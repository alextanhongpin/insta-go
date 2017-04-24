const updateUserButton = document.getElementById('updateUser')
const uploadUserphoto = document.getElementById('uploadUserphoto')
const userphoto = document.getElementById('userphoto')
const username = document.getElementById('username')
const firstName = document.getElementById('firstName')
const lastName = document.getElementById('lastName')

updateUserButton.addEventListener('click', (evt) => {
  evt.preventDefault()
  fetch('http://localhost:8080/users/me', {
    method: 'post',
    headers: {
      'Content-Type': 'application/json; charset=utf-8'
    },
    body: JSON.stringify({
      username: username.value,
      first_name: firstName.value,
      last_name: lastName.value
    }),
    credentials: 'same-origin',
    mode: 'cors'
  }).then((body) => {
    console.log('Body:', body)
    return body.json()
  }).then((json) => {
    console.log('JSON:', json)
  }).catch((err) => {
    console.log('Error:', err)
  })
  return false
})

uploadUserphoto.addEventListener('click', (evt) => {
  evt.preventDefault()
  const formData = new FormData()
  formData.append('file', userphoto.files[0])
  fetch('/users/userphotos', {
    method: 'POST',
    // headers: {
    //  "Content-Type": "multipart/form-data"
    // },
    body: formData,
    mode: 'cors',
    // This is necessary to pass in the cookie
    credentials: 'same-origin'
  }).then((body) => {
    console.log(body)
    return body.json()
  }).then((json) => {
    console.log(json)
  }).catch((err) => {
    console.log(err)
  })
  return false
}, false)

// fetch('/users', {
//   method: 'GET'
// }).then((data) => {
//   return data.json()
// }).then((json) => {
//   console.log(json)
// })
