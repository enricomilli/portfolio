/** @type {import('tailwindcss').Config} */

module.exports = {
    content: ["./**/*.templ", "./**/*.go"],
    theme: {
        extend: {},
    },
    plugins: [require("tailwindcss-motion")],
}
