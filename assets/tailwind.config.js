const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
    purge: [
        '../templates/**/*.tmpl'
    ],
    darkMode: false, // or 'media' or 'class'
    theme: {
        extend: {
            colors: {
                'primary': '#3c3',
                'danger': '#ffdd57',
            },
            zIndex: {
                '-10': '-10',
            },
            fontFamily: {
                sans: ['Raleway', ...defaultTheme.fontFamily.sans],
            },
        },
    },
    variants: {
        extend: {},
    },
    plugins: [],
}
