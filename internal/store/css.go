package store

var CriticalCSS = `
<style>
:root {
  /* Fonts */
  --font-body: Space Mono, monospace;
  --font-title: Montserrat, sans-serif;

  /* Light Theme Colors (Inverted from the website's dark theme) */
  --color-surface: #ffffff;
  --color-surface-alt: #fafafa;
  --color-on-surface: #111111;
  --color-on-surface-strong: #000000;
  --color-primary: #a3e635; /* Primary lime green accent */
  --color-on-primary: #000000;
  --color-secondary: #404040; /* Neutral dark gray */
  --color-on-secondary: #ffffff;
  --color-outline: #d4d4d4;
  --color-outline-strong: #a3a3a3;

  /* Dark Theme Colors (Matched to the website) */
  --color-surface-dark: #0D0C14; /* Main dark background */
  --color-surface-dark-alt: #1a1a1a; /* Slightly lighter surface */
  --color-on-surface-dark: #d1d1d1; /* Primary text color */
  --color-on-surface-dark-strong: #ffffff; /* Stronger text color */
  --color-primary-dark: #6366F1; /* The website's vibrant lime green */
  --color-on-primary-dark: #0a0b0a; /* Text on primary buttons is black */
  --color-secondary-dark: #3B82F6; /* A neutral gray for secondary elements */
  --color-on-secondary-dark: #0a0b0a;
  --color-outline-dark: #2a2a2a; /* Subtle borders */
  --color-outline-dark-strong: #737373; /* Stronger borders */

  /* Shared Colors */
  --color-info: #0284c7;
  --color-on-info: #000000;
  --color-success: #059669;
  --color-on-success: #000000;
  --color-warning: #f59e0b;
  --color-on-warning: #000000;
  --color-danger: #ef4444;
  --color-on-danger: #000000;

  /* Border Radius */
  --radius-none: 0;
  --radius-radius: 0; /* The website uses sharp corners */

  /* Shadows */
  --shadow-sm: 0 1px 2px 0 rgb(0 0 0 / 0.05);
  --shadow-md: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);
  --shadow-lg: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);
  --shadow-xl: 0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1);
}
  </style>
  `
