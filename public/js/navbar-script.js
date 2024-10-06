const navItems = document.getElementById('dropdown-menu');
const iconOpen = document.getElementById('icon-open');
const iconClose = document.getElementById('icon-close');
const openNavBtn = document.getElementById('menu-button')

// Initial setup
if (window.innerWidth >= 768) {
    desktopNavState()
} else {
    mobileClosedMenuState()
}

// Handle window resize
window.addEventListener('resize', function() {
    if (window.innerWidth >= 768) {
        desktopNavState()
    } else {
        mobileClosedMenuState()
    }
});

iconOpen.addEventListener('click', ()=> {
    mobileOpenMenuState()
})

iconClose.addEventListener('click', ()=> {
    mobileClosedMenuState()
})

function mobileOpenMenuState() {
    show(openNavBtn)
    hide(iconOpen)
    show(navItems)
    show(iconClose)
}

function mobileClosedMenuState() {
    show(openNavBtn)
    hide(navItems)
    hide(iconClose)
    show(iconOpen)
}

// makes function global
window.mobileClosedMenuState = mobileClosedMenuState;

function desktopNavState() {
    show(navItems)
    hide(iconOpen)
    hide(iconClose)
    hide(openNavBtn)
}

function show(htmlElem) {
    htmlElem.classList.add('flex')
    htmlElem.classList.remove('hidden')
}

function hide(htmlElem) {
    htmlElem.classList.add('hidden')
    htmlElem.classList.remove('flex')
}