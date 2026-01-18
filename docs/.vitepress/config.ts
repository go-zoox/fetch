import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Fetch',
  description: 'HTTP Client for Go, inspired by the Fetch API',
  base: '/fetch/',
  lang: 'en-US',
  
  locales: {
    root: {
      label: 'English',
      lang: 'en',
      link: '/',
      title: 'Fetch',
      description: 'HTTP Client for Go, inspired by the Fetch API',
      themeConfig: {
        nav: [
          { text: 'Home', link: '/' },
          { text: 'Quick Start', link: '/quickstart' },
          { text: 'Guide', link: '/guide/' },
          { text: 'API Reference', link: '/api/' },
          { text: 'Examples', link: '/examples/' },
          { text: 'GitHub', link: 'https://github.com/go-zoox/fetch' }
        ],
        sidebar: {
          '/guide/': [
            {
              text: 'Guide',
              items: [
                { text: 'Installation', link: '/guide/installation' },
                { text: 'Configuration', link: '/guide/config' },
                { text: 'HTTP Methods', link: '/guide/methods' },
                { text: 'Response Handling', link: '/guide/response' },
                { text: 'Timeout & Retry', link: '/guide/timeout' },
                { text: 'Proxy', link: '/guide/proxy' },
                { text: 'Authentication', link: '/guide/auth' },
                { text: 'Upload & Download', link: '/guide/upload-download' },
                { text: 'Advanced', link: '/guide/advanced' }
              ]
            }
          ],
          '/api/': [
            {
              text: 'API Reference',
              items: [
                { text: 'Overview', link: '/api/' },
                { text: 'Global Functions', link: '/api/globals' },
                { text: 'Fetch', link: '/api/fetch' },
                { text: 'Config', link: '/api/config' },
                { text: 'Response', link: '/api/response' },
                { text: 'Methods', link: '/api/methods' },
                { text: 'File Operations', link: '/api/file-ops' }
              ]
            }
          ],
          '/examples/': [
            {
              text: 'Examples',
              items: [
                { text: 'Overview', link: '/examples/' },
                { text: 'Basic Usage', link: '/examples/basic' },
                { text: 'HTTP Methods', link: '/examples/http-methods' },
                { text: 'Authentication', link: '/examples/auth' },
                { text: 'File Operations', link: '/examples/file-operations' },
                { text: 'Timeout & Retry', link: '/examples/timeout-retry' },
                { text: 'Proxy', link: '/examples/proxy' },
                { text: 'Stream', link: '/examples/stream' },
                { text: 'Session & Cookies', link: '/examples/session-cookies' },
                { text: 'Context Cancellation', link: '/examples/context-cancel' },
                { text: 'Error Handling', link: '/examples/error-handling' }
              ]
            }
          ]
        },
        socialLinks: [
          { icon: 'github', link: 'https://github.com/go-zoox/fetch' }
        ],
        search: {
          provider: 'local'
        }
      }
    },
    zh: {
      label: '中文',
      lang: 'zh-CN',
      link: '/zh/',
      title: 'Fetch',
      description: 'Go 语言 HTTP 客户端库，受 Fetch API 启发',
      themeConfig: {
        nav: [
          { text: '首页', link: '/zh/' },
          { text: '快速开始', link: '/zh/quickstart' },
          { text: '指南', link: '/zh/guide/' },
          { text: 'API 参考', link: '/zh/api/' },
          { text: '示例', link: '/zh/examples/' },
          { text: 'GitHub', link: 'https://github.com/go-zoox/fetch' }
        ],
        sidebar: {
          '/zh/guide/': [
            {
              text: '指南',
              items: [
                { text: '安装', link: '/zh/guide/installation' },
                { text: '配置', link: '/zh/guide/config' },
                { text: 'HTTP 方法', link: '/zh/guide/methods' },
                { text: '响应处理', link: '/zh/guide/response' },
                { text: '超时与重试', link: '/zh/guide/timeout' },
                { text: '代理', link: '/zh/guide/proxy' },
                { text: '认证', link: '/zh/guide/auth' },
                { text: '上传与下载', link: '/zh/guide/upload-download' },
                { text: '高级特性', link: '/zh/guide/advanced' }
              ]
            }
          ],
          '/zh/api/': [
            {
              text: 'API 参考',
              items: [
                { text: '概览', link: '/zh/api/' },
                { text: '全局函数', link: '/zh/api/globals' },
                { text: 'Fetch', link: '/zh/api/fetch' },
                { text: 'Config', link: '/zh/api/config' },
                { text: 'Response', link: '/zh/api/response' },
                { text: 'Methods', link: '/zh/api/methods' },
                { text: '文件操作', link: '/zh/api/file-ops' }
              ]
            }
          ],
          '/zh/examples/': [
            {
              text: '示例',
              items: [
                { text: '概览', link: '/zh/examples/' },
                { text: '基础用法', link: '/zh/examples/basic' },
                { text: 'HTTP 方法', link: '/zh/examples/http-methods' },
                { text: '认证', link: '/zh/examples/auth' },
                { text: '文件操作', link: '/zh/examples/file-operations' },
                { text: '超时与重试', link: '/zh/examples/timeout-retry' },
                { text: '代理', link: '/zh/examples/proxy' },
                { text: '流式传输', link: '/zh/examples/stream' },
                { text: '会话与 Cookie', link: '/zh/examples/session-cookies' },
                { text: 'Context 取消', link: '/zh/examples/context-cancel' },
                { text: '错误处理', link: '/zh/examples/error-handling' }
              ]
            }
          ]
        },
        socialLinks: [
          { icon: 'github', link: 'https://github.com/go-zoox/fetch' }
        ],
        search: {
          provider: 'local'
        }
      }
    }
  },

  themeConfig: {
    // logo: '/logo.svg', // Uncomment when logo is added
  }
})
