const searchGoogle = document.querySelector("#submit-terms")
searchGoogle.addEventListener("submit", function(e) {
  e.preventDefault()

  if (String(termsContent).trim() == "") {
    alert("É necessário ao menos um termo para realizar a pesquisa")
    return
  }

  if (String(email).trim() == "") {
    alert("Adicione um e-mail para envio do reporte")
    return
  }

  fetchResult()

  window.location.assign("thanks.html")
})

function fetchResult() {
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

  apiURL = ""

  const environment = window.location.hostname

  if (environment != "127.0.0.1") {
    apiURL = "https://brand-monitor-gzbnmirggq-ue.a.run.app"
  } else {
    apiURL = "http://localhost:3000"
  }

  fetch(`${apiURL}/search?location=Brazil`, opt).
  then(res => {
    if (res.status == 404) {
      alert("Nenhum resultado encontrado para sua busca nas páginas 1 - 3 do Google")
      return
    }

    if (res.status != 200) {
      resInJson = res.json()
      alert(resInJson)
      return
    }

    alert("E-mail enviado para seu inbox!")
  }).
  catch(err => err.json()).
  catch(data => {
    alert(data)
  })
}