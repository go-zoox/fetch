# Fetch Documentation

This directory contains the VitePress documentation site for go-zoox/fetch.

## Development

To start the development server:

```bash
cd docs
npm install
npm run dev
```

## Build

To build the documentation:

```bash
cd docs
npm run build
```

The built files will be in `docs/.vitepress/dist/`.

## Structure

- `index.md` - English homepage
- `quickstart.md` - Quick start guide (English)
- `guide/` - Detailed guides (English)
- `api/` - API reference (English)
- `examples/` - Code examples (English)
- `zh/` - Chinese documentation
  - `index.md` - Chinese homepage
  - `quickstart.md` - Quick start guide (Chinese)
  - `guide/` - Detailed guides (Chinese)
  - `api/` - API reference (Chinese)
  - `examples/` - Code examples (Chinese)
- `.vitepress/` - VitePress configuration
  - `config.ts` - Main configuration file
  - `public/` - Static assets

## Deployment

The documentation is automatically deployed to GitHub Pages via GitHub Actions when changes are pushed to the `master` branch.
