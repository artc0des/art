const btn = document.getElementById('menu-btn')
const menu = document.getElementById('menu')
const header = document.getElementById('header')
const intro = document.getElementById('hero')

if (intro) {
    setTimeout(function() {
        intro.classList.remove('opacity-0')
        intro.classList.add('opacity-100')
    }, 500) 
}
btn.addEventListener('click', navToggle)
function navToggle() {
    btn.classList.toggle('open')
    menu.classList.toggle('flex')
    menu.classList.toggle('hidden')
}

