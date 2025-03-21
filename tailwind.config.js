/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./ui/**/*.{html,js}"],
  theme: {
    screens: {
      sm: '480px',
      md: '768px',
      lg: '976px',
      xl: '1440px'
    },
    fontFamily: {
      alma: ["Almarai", "sans-serif"],
    },
    extend: {
      colors: {
        bgBlue: 'hsl(212.31, 17.81%, 28.63%)',
        titleGray: 'hsl(167.14, 38.89%, 92.94%)',
        contentWhite: 'hsl(205, 100%, 97.65%)',
        borderGold: 'hsl(37.12, 50.43%, 54.12%)',
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}

