const form = document.getElementById("form")
const inputFile = document.getElementById("file")
const inputValue = document.getElementById("value")
const notification = document.getElementById("apiResponse")


const handleSubmit = async (event) => {
	event.preventDefault()
	const formData = new FormData()


	for (let i = 0; i < inputFile.files.length; i++) {
		formData.append("upload", inputFile.files[i])
	}
	formData.append("airac", inputValue.value)

	await fetch("api/v1/task", {
		method: "post",
		body: formData
	})
		.then(response => response.json())
		.then(json => {
			if (json["message"] !== undefined) {
				inputValue.value = ""
				inputFile.value = ""
				notification.style.background = "green"
				notification.innerHTML = JSON.stringify(json["message"])
				notification.style.display = "block"
			} else {
				notification.style.background = "red"
				notification.innerHTML = JSON.stringify(json["msg"])
				notification.style.display = "block"
			}
		})
		.catch(_ => {
			notification.style.background = "red"
			notification.innerHTML = "Something wrong"
			notification.style.display = "block"
		});
}

form.addEventListener("submit", handleSubmit)
