
// Sample example on how to upload the photo
console.log('loaded')
// Yeah, poorly set global
let fileToUpload = null

const uploader = document.getElementById('uploader')
const submit = document.getElementById('submit')
const caption = document.getElementById('caption')

uploader.addEventListener('change', handleFileChange, false)

function handleFileChange (evt) {
  const files = evt.currentTarget.files
  const file = files[0]
  const reader = new FileReader()

  reader.onload = function (e) {
    fileToUpload = e.currentTarget.result
  }
  reader.readAsText(file)
}

const URL = '/api/photos'
submit.addEventListener('click', (evt) => {
  evt.preventDefault()
  const formData = new FormData()
  formData.append('file', uploader.files[0])
  formData.append('caption', caption.value)
  console.log(uploader.files[0])
  fetch(URL, {
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

function getPhotoCount () {
  fetch('/api/photos/count', {
    credentials: 'same-origin'
  })
  .then((body) => body.json())
  .then((json) => {
    console.log(json)
  }).catch((err) => {
    console.log(err)
  })
}

function getPhotos () {
  return fetch('/api/photos', {
    credentials: 'same-origin'
  })
  .then(body => body.json())
  .catch((error) => {
    console.log(error)
  })
}

function photosView () {
  getPhotos().then((photos) => {
    console.log('get photos', photos)
    const photoView = document.getElementById('photoView')
    photoView.innerHTML = ''
    const photosView = photos.data.map((photo) => {
      console.log(photo)
      return `
        <div>
          <img src="${photo.src}" width="400" height="auto" alt="${photo.caption}"/>
          <div>${photo.caption}</div>
          <div>Like Count: ${photo.like_count || 0}</div>
          <div>${photo.user_likes.map((user) => `<div>${user}</div>`).join(',')}</div>
          <button>Like</button>
        </div>
      `
    })
    photoView.innerHTML += photosView.join('')
  })
}
photosView()
