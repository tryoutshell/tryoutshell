// @ts-check
import { themes as prismThemes } from "prism-react-renderer";

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: "Tryoutshell",
  tagline: "Learn by doing — directly in your terminal",
  favicon: "/img/favicon.ico",

  url: "https://tryoutshell.dev", // your domain
  baseUrl: "/",

  organizationName: "tryoutshell", // GitHub org/user
  projectName: "tryoutshell", // GitHub repo

  onBrokenLinks: "warn",
  onBrokenMarkdownLinks: "warn",

  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  presets: [
    [
      "classic",
      {
        docs: {
          path: "../docs",
          routeBasePath: "/docs",
          sidebarPath: "./sidebars.js",
          include: ["**/*.{md,mdx}"],
        },
        blog: false,
        theme: { customCss: "./src/css/custom.css" },
      },
    ],
  ],

  themeConfig: {
    image: "img/tryoutshell-og.png",
    navbar: {
      title: "Tryoutshell",
      logo: {
        alt: "Tryoutshell Logo",
        src: "img/logo.png",
      },
      items: [
        { type: "search", position: "right" },
        {
          type: "doc",
          docId: "intro",
          label: "Introduction",
          position: "left",
        },
        {
          type: "doc",
          docId: "getting-started/index",
          label: "Getting Started",
          position: "left",
        },
        {
          type: "doc",
          docId: "examples/step-by-step-tutorial",
          label: "Examples",
          position: "left",
        },
        {
          href: "https://github.com/tryoutshell/tryoutshell",
          position: "right",
          className: "header-github-link",
          "aria-label": "GitHub repository",
        },
      ],
    },
    footer: {
      style: "dark",
      links: [
        {
          title: "Docs",
          items: [
            { label: "Quick Start", to: "/docs/getting-started/quick-start" },
          ],
        },
        {
          title: "Community",
          items: [
            {
              label: "GitHub",
              href: "https://github.com/tryoutshell/tryoutshell",
            },
            { label: "Twitter", href: "https://x.com/tryoutshell" },
          ],
        },
      ],
      copyright: `© ${new Date().getFullYear()} Tryoutshell Contributors. Built with Docusaurus.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    },
  },
};

export default config;
