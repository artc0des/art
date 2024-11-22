/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./*.{html,js}"],
  theme: {
    screens: {
      sm: '480px',
      md: '768px',
      lg: '976px',
      xl: '1440px'
    },
    extend: {
      colors: {
        oliveGreen: 'hsl(127.5, 83.33%, 90.59%)',
        seaCreatureGreen: 'hsl(168.73, 100%, 38.63%)',
        charcoalBG: 'hsl(190, 6.52%, 18.04%)',
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}

