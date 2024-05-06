const addTermForm = document.querySelector("#add-term")

addTermForm.addEventListener("submit", function (e) {
  e.preventDefault()

  const term = document.querySelector("#input")
  const termText = String(term.value).trim()
  const TermsList = document.querySelector("#terms-list")

  if (termText == "") {
    alert("Campo de termos não pode estar vazio")
    return
  }

  if (termsContent != "") {
    termsContent += `,${termText}`
  } else {
    termsContent += `${termText}`
  }

  TermsList.className = "terms-list"
  TermsList.innerHTML = `<p>${termsContent}</p>`
})