import { animate } from "motion";

export function animateRandomNum() {
  const num = document.getElementById("random-number");

 animate(num, { opacity: 1, transform: "rotate(720deg)"}, {duration: 0.6});
}


window.animations = window.animations || {};
window.animations.RandNum = animateRandomNum