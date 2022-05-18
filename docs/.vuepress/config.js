module.exports = {
  theme: 'cosmos',
  title: 'Enron Documentation',
  locales: {
    '/': {
      lang: 'en-US'
    },
  },
  markdown: {
    extendMarkdown: (md) => {
      md.use(require("markdown-it-katex"));
    },
  },
  head: [
    [
      "link",
      {
        rel: "stylesheet",
        href:
          "https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.5.1/katex.min.css",
      },
    ],
    [
      "link",
      {
        rel: "stylesheet",
        href:
          "https://cdn.jsdelivr.net/github-markdown-css/2.2.1/github-markdown.css",
      },
    ],
  ],
  base: process.env.VUEPRESS_BASE || '/',
  plugins: [
    'vuepress-plugin-element-tabs'
  ],
  head: [
    // ['link', { rel: "apple-touch-icon", sizes: "180x180", href: "/apple-touch-icon.png" }],
    ['link', { rel: "icon", type: "image/png", sizes: "32x32", href: "/favicon32.png" }],
    ['link', { rel: "icon", type: "image/png", sizes: "16x16", href: "/favicon16.png" }],
    ['link', { rel: "manifest", href: "/site.webmanifest" }],
    ['meta', { name: "msapplication-TileColor", content: "#2e3148" }],
    ['meta', { name: "theme-color", content: "#ffffff" }],
    ['link', { rel: "icon", type: "image/svg+xml", href: "/favicon.svg" }],
    // ['link', { rel: "apple-touch-icon-precomposed", href: "/apple-touch-icon-precomposed.png" }],
  ],
  themeConfig: {
    repo: 'tharsis/enron',
    docsRepo: 'tharsis/enron',
    docsBranch: 'main',
    docsDir: 'docs',
    editLinks: true,
    custom: true,
    project: {
      name: 'Enron',
      denom: 'Enron',
      ticker: 'ECH',
      binary: 'enrond',
      testnet_denom: 'tEnron',
      testnet_ticker: 'tENRON',
      rpc_url: 'https://eth.bd.enron.org:8545',
      rpc_url_testnet: 'https://eth.bd.enron.dev:8545',
      rpc_url_local: 'http://localhost:8545/',
      chain_id: '3000',
      testnet_chain_id: '9003',
      latest_version: 'v1.0.0-Beta',
      version_number: '2',
      testnet_version_number: '3',
      testnet_evm_explorer_url: 'https://evm.enron.dev',
      evm_explorer_url: 'https://evm.enron.org',
      testnet_cosmos_explorer_url: 'https://explorer.enron.dev/',
      cosmos_explorer_url: 'https://www.mintscan.io/enron',
    },
    logo: {
      src: '/enron-black.svg',
    },
    algolia: {
      id: 'BH4D9OD16A',
      key: 'a5d55fe5f540cc3bd28fa2c72f2b5bd8',
      index: 'enron'
    },
    topbar: {
      banner: false
    },
    sidebar: {
      auto: false,
      nav: [
        {
          title: 'Reference',
          children: [
            {
              title: 'Introduction',
              directory: true,
              path: '/intro'
            },
            {
              title: 'Quick Start',
              directory: true,
              path: '/quickstart'
            },
            {
              title: 'Basics',
              directory: true,
              path: '/basics'
            },
            {
              title: 'Core Concepts',
              directory: true,
              path: '/core'
            },
          ]
        },
        {
          title: 'Governance',
          children: [
            {
              title: 'Enron Governance',
              directory: true,
              path: '/governance'
            },
          ]
        },
        {
          title: 'Guides',
          children: [
            {
              title: 'Localnet',
              directory: true,
              path: '/guides/localnet'
            },
            {
              title: 'Keys and Wallets',
              directory: true,
              path: '/guides/keys-wallets'
            },
            {
              title: 'Ethereum Tooling',
              directory: true,
              path: '/guides/tools'
            },
            {
              title: 'Validators',
              directory: true,
              path: '/guides/validators'
            },
            {
              title: 'Upgrades',
              directory: true,
              path: '/guides/upgrades'
            },
            {
              title: 'Key Management System',
              directory: true,
              path: '/guides/kms'
            },
          ]
        },
        {
          title: 'APIs',
          children: [
            {
              title: 'JSON-RPC',
              directory: true,
              path: '/api/json-rpc'
            },
            {
              title: 'Protobuf Reference',
              directory: false,
              path: '/api/proto-docs'
            },
          ]
        },
        // {
        //   title: 'Clients',
        //   children: [
        //     {
        //       title: 'APIs',
        //       directory: false,
        //       path: '/clients/apis'
        //     },
        //     {
        //       title: 'Enronjs',
        //       directory: false,
        //       path: '/clients/enronjs'
        //     },
        //   ]
        // },
        {
          title: 'Mainnet',
          children: [
            {
              title: 'Join Mainnet',
              directory: false,
              path: '/mainnet/join'
            },
          ]
        },
        {
          title: 'Testnet',
          children: [
            {
              title: 'Join Testnet',
              directory: false,
              path: '/testnet/join'
            },
            {
              title: 'Token Faucet',
              directory: false,
              path: '/testnet/faucet'
            },
            {
              title: 'Deploy Node on Cloud',
              directory: false,
              path: '/testnet/cloud_providers'
            }
          ]
        },
        {
          title: 'Specifications',
          children: [{
            title: 'Modules',
            directory: true,
            path: '/modules'
          }]
        },
        {
          title: 'Block Explorers',
          children: [
            {
              title: 'Block Explorers',
              path: '/tools/explorers'
            },
            {
              title: 'Blockscout (EVM)',
              path: 'https://evm.enron.org'
            },
            {
              title: 'Mintscan (Cosmos)',
              path: 'https://www.mintscan.io/enron/'
            },
          ]
        },
        {
          title: 'Ecosystem',
          children: [
            {
              title: 'Awesome Enron',
              path: 'https://github.com/tharsis/awesome'
            },
            {
              title: 'Enron Space',
              path: 'https://enron.space/'
            }
          ]
        },
        {
          title: 'Resources',
          children: [
            {
              title: 'Enron Go API',
              path: 'https://pkg.go.dev/github.com/enron/enron'
            },
            {
              title: 'Ethermint Library Go API',
              path: 'https://pkg.go.dev/github.com/tharsis/ethermint'
            },
            {
              title: 'Enron gRPC Gateway API',
              path: 'https://api.enron.dev/'
            },
            {
              title: 'JSON-RPC API',
              path: '/api/json-rpc/endpoints'
            }
          ]
        }
      ]
    },
    gutter: {
      title: 'Help & Support',
      chat: {
        title: 'Developer Chat',
        text: 'Chat with Enron developers on Discord.',
        url: 'https://discord.gg/enron',
        bg: 'linear-gradient(103.75deg, #1B1E36 0%, #22253F 100%)'
      },
      forum: {
        title: 'Enron Developer Forum',
        text: 'Join the Enron Developer Forum to learn more.',
        url: 'https://forum.cosmos.network/c/ethermint', // TODO: replace with commonwealth link
        bg: 'linear-gradient(221.79deg, #3D6B99 -1.08%, #336699 95.88%)',
        logo: 'ethereum-white'
      },
      github: {
        title: 'Found an Issue?',
        text: 'Help us improve this page by suggesting edits on GitHub.',
        bg: '#F8F9FC'
      }
    },
    footer: {
      logo: '/enron-black.svg',
      textLink: {
        text: 'enron.org',
        url: 'https://enron.org'
      },
      services: [
        {
          service: 'github',
          url: 'https://github.com/enron/enron'
        },
        {
          service: "twitter",
          url: "https://twitter.com/EnronOrg",
        },
        {
          service: "telegram",
          url: "https://t.me/EnronOrg",
        },
        {
          service: "linkedin",
          url: "https://www.linkedin.com/company/tharsis-finance/",
        },
        {
          service: "medium",
          url: "https://enron.blog/",
        },
      ],
      smallprint: 'This website is maintained by Tharsis Labs Ltd.',
      links: [{
        title: 'Documentation',
        children: [{
          title: 'Cosmos SDK Docs',
          url: 'https://docs.cosmos.network/master/'
        },
        {
          title: 'Ethereum Docs',
          url: 'https://ethereum.org/developers'
        },
        {
          title: 'Tendermint Core Docs',
          url: 'https://docs.tendermint.com'
        }
        ]
      },
      {
        title: 'Community',
        children: [{
          title: 'Enron Community',
          url: 'https://discord.gg/enron'
        },
        ]
      },
      {
        title: 'Tharsis',
        children: [
          {
            title: 'Jobs at Tharsis',
            url: 'https://tharsis.notion.site/'
          }
        ]
      }
      ]
    },
    versions: [
      {
        "label": "main",
        "key": "main"
      },
    ],
  }
};
