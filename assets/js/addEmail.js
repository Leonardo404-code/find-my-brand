const addEmailForm = document.querySelector("#add-email")

addEmailForm.addEventListener("submit", function(e) {
  e.preventDefault()
  const emailInput = document.querySelector("#email")
  const emailText = String(emailInput.value)
  if (emailText.trim() == "") {
    alert("Campo de email não pode estar vazio")
    return
  }

  email = String(emailInput.value)
  alert("e-mail adicionado")
})