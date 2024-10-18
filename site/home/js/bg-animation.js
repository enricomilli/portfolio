import { animate } from "motion";

function animateBg(id) {
    const mainBody = document.getElementById(id)
    console.log("called on", id)
    const shapes = ['circle', 'square', 'triangle'];
    const colors = ['#FF6B6B', '#4ECDC4', '#45B7D1', '#FFA07A'];

    for (let i = 0; i < 30; i++) {
        const shape = document.createElement('div');
        shape.className = `floating-shape ${shapes[Math.floor(Math.random() * shapes.length)]}`;
        shape.style.left = `${Math.random() * 100}%`;
        shape.style.top = `${Math.random() * 100}%`;
        shape.style.borderColor = colors[Math.floor(Math.random() * colors.length)];
        shape.style.borderWidth = "2px"
        mainBody.appendChild(shape);

        animateShape(shape);
    }
}

function animateShape(shape) {
    const randomX = () => `${(Math.random() - 0.5) * 300}%`;
    const randomY = () => `${(Math.random() + 0.5) * 300}%`;
    const randomRotate = () => Math.random() * 720 - 360;
    const randomScale = () => 0.2 + Math.random() * 0.5;

    animate(
        shape,
        {
            x: [randomX(), randomX(), randomX()],
            y: [randomY(), randomY(), randomY()],
            rotate: [randomRotate(), randomRotate(), randomRotate()],
            scale: [randomScale(), randomScale(), randomScale()],
        },
        {
            duration: 15 + Math.random() * 15,
            repeat: Infinity,
            easing: "ease-in-out",
            direction: "alternate",
            onComplete: () => animateShape(shape),
        }
    );
}

window.animations = window.animations || {};
window.animations.Background = animateBg;
