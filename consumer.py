import time
import json
import shutil
import urllib.request
import os
import redis
import requests


REDIS_HOST, REDIS_PORT, REDIS_DB, REDIS_QUEUE_NAME = "localhost", 6379, 1, "rmq::queue::[tasks]::ready"
THROAT_SECONDS = 1
DATA_STORAGE_HOST = "http://localhost:3000"
TEMP_STORAGE = "./temp"
RM_LIMIT = 10737418240


class Consumer:
    def __init__(self):
        self.connection = None
        self.host = REDIS_HOST
        self.port = REDIS_PORT
        self.db = REDIS_DB
        self.queue_name = REDIS_QUEUE_NAME
        self.connect()

    def connect(self):
        connection = redis.Redis(host=self.host, port=self.port, db=self.db)
        connection.ping()
        self.connection = connection

    def listen(self):
        while True:
            queue_length = self.connection.llen(self.queue_name)
            if queue_length >= 0:
                task = self.connection.brpop(self.queue_name)
                if task:
                    Task(task).run()
                    pass
            time.sleep(THROAT_SECONDS)


class Task:
    def __init__(self, message_string):
        self.message_string = message_string
        self.uuid = None
        self.publishedAt = None
        self.airac = None
        self.files = None
        self.files_paths = None
        self.status = "PROCESSING"
        self.prepare_message_string()

    def prepare_message_string(self):
        data = self.message_string[1]
        data = json.loads(data.decode("utf-8"))
        self.uuid = data.get("Uuid")
        self.publishedAt = data.get("PublishedAt")
        self.airac = data.get("Airac")
        self.files = data.get("Files")

    def run(self):
        self.update_db_task_status()

        try:
            self.download()
            pass
        except Exception as error:
            self.status = "ERROR_IN_DOWNLOAD"
            self.update_db_task_status()
            return error

        try:
            self.run_script()
            pass
        except Exception as error:
            self.status = "ERROR_IN_PROCESS"
            self.update_db_task_status()
            return error

        self.status = "SUCCESS"
        self.update_db_task_status()

    def download(self):
        self.clean_data()
        self.files_paths = []
        for file in self.files:
            filepath = f"{TEMP_STORAGE}/{file}"
            self.files_paths.append(filepath)
            if not os.path.isfile(filepath):
                url = f"{DATA_STORAGE_HOST}/uploads/{file}"
                urllib.request.urlretrieve(url, filepath)

    def run_script(self):
        print("---   ---   ---")
        print("START PROCESSING")
        print(self.airac, self.files_paths)
        # RUN SCRIPT HERE
        time.sleep(10)
        print("---   ---   ---\n\n")

    def update_db_task_status(self):
        url = f"{DATA_STORAGE_HOST}/api/v1/task"
        data = {
            "uuid": self.uuid,
            "status": self.status
        }
        requests.put(url, json=data)

    @staticmethod
    def clean_data():
        if not os.path.exists(TEMP_STORAGE):
            os.makedirs(TEMP_STORAGE)
        size = sum(d.stat().st_size for d in os.scandir(TEMP_STORAGE) if d.is_file())
        if size > RM_LIMIT:
            shutil.rmtree(TEMP_STORAGE)


if __name__ == "__main__":
    Consumer().listen()
