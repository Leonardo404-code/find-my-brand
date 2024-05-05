const addTermForm = document.querySelector("#add-term")

termsContent = ""

addTermForm.addEventListener("submit", function (e) {
  e.preventDefault()

  const term = document.querySelector("#input")
  const termText = String(term.value)
  const TermsList = document.querySelector("#terms-list")

  if (termsContent != "") {
    termsContent += `,${termText}`
  } else {
    termsContent += `${termText}`
  }

  TermsList.className = "terms-list"
  TermsList.innerHTML = `<p>${termsContent}</p>`
})

const searchGoogleForm = document.querySelector("#submit-terms")

searchGoogleForm.addEventListener("submit", function(e) {
  e.preventDefault()

  const opt = {
    method: "POST",
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      terms: termsContent
    })
  }

  fetch("http://localhost:3000?location=Brazil", opt).
  then(res => {
    if (res.status == 400) {
      alert("Nenhum resultado encontrado para sua busca nas páginas 1 - 5 do Google")
    }

    alert("E-mail enviado para seu inbox!")
  }).
  catch(err => {
    alert(err)
  })
})