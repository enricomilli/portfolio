import { animate } from "motion";

function animateHero() {
    // const heroSection = document.getElementById('hero-section');
    const heroBrand = document.getElementById('hero-brand');
    const heroTitle = document.getElementById('hero-title');
    const heroDesc = document.getElementById('hero-desc');
    const heroButton = document.getElementById('hero-button');

    // Animate hero brand
    animate(heroBrand, 
        { opacity: [0, 1], y: [20, 0] },
        { duration: 0.8, easing: "ease-out" }
    );

    // Animate hero title
    animate(heroTitle, 
        { opacity: [0, 1], y: [30, 0] },
        { duration: 0.8, delay: 0.4, easing: "ease-out" }
    );

    // Animate hero button 
    animate(heroButton, 
        { opacity: [0, 1], y: [30, 0] },
        { duration: 0.8, delay: 1, easing: "ease-out" }
    );

    animate(heroDesc, 
        { opacity: [0, 1], y: [30, 0] },
        { duration: 0.8, delay: 0.5, easing: "ease-out" }
    );

    // Typewriter effect for hero title
    typewriterEffect(heroDesc, 10);

    // // Parallax effect on scroll
    // const mainBody = document.getElementById('main')
    // window.addEventListener('scroll', () => {
    //     const scrollY = window.scrollY;
    //     animate(mainBody, { y: scrollY * 0.5 }, { duration: 0 });
    // });
}

function typewriterEffect(element, speed) {
    const text = element.innerText;
    element.innerHTML = '';
    let i = 0;
    const timer = setInterval(() => {
        if (i < text.length) {
            element.innerHTML += text.charAt(i);
            i++;
        } else {
            clearInterval(timer);
        }
    }, speed);
}


window.animations = window.animations || {};
window.animations.Hero = animateHero