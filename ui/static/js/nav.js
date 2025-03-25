const btn = document.getElementById('menu-btn')
const menu = document.getElementById('menu')
const header = document.getElementById('header')
const intro = document.getElementById('hero')
const home = document.getElementById('home_link')
const about = document.getElementById('about_link')
const experience = document.getElementById('experience_link')
const contact = document.getElementById('contact_link')
const projects = document.getElementById('projects') 

const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
        console.log(entry)
        if (entry.isIntersecting) {
            entry.target.classList.add('show');
            entry.target.classList.remove('hide')
        } else {
            entry.target.classList.remove('show')
            entry.target.classList.add('hide');
        }
    });
});

const hiddenElements = document.querySelectorAll('.hide')
hiddenElements.forEach((el) => observer.observe(el))

window.onload = function() {
    let path = window.location.href
    let url = new URL(path)
    highlight(url.pathname)
};

function highlight(url) {
    console.log("received", url)
    if (url === "/") {
        home.classList.add('border-b-2')
        home.classList.add('border-b-borderGold')
        about.classList.remove('border-b-2')
        about.classList.remove('border-b-borderGold')
        experience.classList.remove('border-b-2')
        experience.classList.remove('border-b-borderGold')
        contact.classList.remove('border-b-2')
        contact.classList.remove('border-b-borderGold')
    }
    else if (url === "/about") {
        home.classList.remove('border-b-2')
        home.classList.remove('border-b-borderGold')
        about.classList.add('border-b-2')
        about.classList.add('border-b-borderGold')
        experience.classList.remove('border-b-2')
        experience.classList.remove('border-b-borderGold')
        contact.classList.remove('border-b-2')
        contact.classList.remove('border-b-borderGold')
    }
    else if (url === "/experience") {
        home.classList.remove('border-b-2')
        home.classList.remove('border-b-borderGold')
        about.classList.remove('border-b-borderGold')
        about.classList.remove('border-b-2')
        experience.classList.add('border-b-borderGold')
        experience.classList.add('border-b-2')
        contact.classList.remove('border-b-borderGold')
        contact.classList.remove('border-b-2')
    }
    else if (url === "/contact") {
        home.classList.remove('border-b-2')
        home.classList.remove('border-b-borderGold')
        about.classList.remove('border-b-borderGold')
        about.classList.remove('border-b-2')
        experience.classList.remove('border-b-borderGold')
        experience.classList.remove('border-b-2')
        contact.classList.add('border-b-2')
        contact.classList.add('border-b-borderGold')
    }
}


if (intro) {
    setTimeout(function() {
        intro.classList.remove('opacity-0')
        intro.classList.add('opacity-100')
    }, 250) 
}

btn.addEventListener('click', navToggle)
function navToggle() {
    btn.classList.toggle('open')
    menu.classList.toggle('flex')
    menu.classList.toggle('hidden')
}

