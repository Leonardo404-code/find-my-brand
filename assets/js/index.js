const addTermForm = document.querySelector("#add-term")

termsContent = ""
email = ""

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

const addEmailForm = document.querySelector("#add-email")
addEmailForm.addEventListener("submit", function(e) {
  e.preventDefault()
  const emailInput = document.querySelector("#email")
  if (String(emailInput).trim() == "") {
    alert("Campo de email não pode estar vazio")
    return
  }

  email = String(emailInput.value)
})

const searchGoogleForm = document.querySelector("#submit-terms")
searchGoogleForm.addEventListener("submit", function(e) {
  e.preventDefault()

  if (String(termsContent).trim() == "") {
    alert("É necessário ao menos um termo para realizar a pesquisa")
    return
  }

  if (String(email).trim() == "") {
    alert("Adicione um e-mail para envio do reporte")
    return
  }

  const opt = {
    method: "POST",
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      terms: termsContent,
      email: email
    })
  }

  fetch("http://localhost:3000?location=Brazil", opt).
  then(res => {
    if (res.status == 400) {
      alert("Nenhum resultado encontrado para sua busca nas páginas 1 - 3 do Google")
      return
    }

    alert("E-mail enviado para seu inbox!")
  }).
  catch(err => {
    alert(err)
  })
})