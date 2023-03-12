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

	fetch("api/v1/task", {
		method: "post",
		body: formData
	}).then((response) => {
		if (response.ok) {
			return response
		}
		throw new Error("Something went wrong")
	})
		.then((apiResponse) => {
			if (apiResponse.code === 200) {
				inputValue.value = ""
				inputFile.value = ""
				notification.style.background = "green"
				notification.innerHTML = JSON.stringify(apiResponse.json()["message"])
				notification.style.display = "block"
			} else {
				notification.style.background = "red"
				notification.innerHTML = JSON.stringify(apiResponse.json()["msg"])
				notification.style.display = "block"
			}
		})
		.catch((error) => {
			notification.style.background = "red"
			notification.innerHTML = error
			notification.style.display = "block"
		})
}

form.addEventListener("submit", handleSubmit)
