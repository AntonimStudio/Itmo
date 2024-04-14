<div align="center">
  <h1 align="center">LeLa</h1>
  <h3>DaPoCa inc.</h3>
</div>

## Features

- **Shareable Links:** Share your document securely by sending a custom link
- **Custom Branding:** Add a custom domain and your own branding
- **Analytics:** Get insights via document tracking and soon page-by-page analytics
- **Self-hosted, open-source:** Host it yourself and hack on it

## Tech Stack

- Golang – Server
- Figma – HTML
- Figma – CSS
- Canva – Presentation
- PostgreSQL – Database

## Getting Started

### Prerequisites

Here's what you need to be able to run LeLa:

- Docker
- PostgreSQL Database

### 1. Clone the repository

```shell
git clone https://github.com/mfts/papermark.git
cd papermark
```

### 2. Install npm dependencies

```shell
npm install
```

### 3. Copy the environment variables to `.env` and change the values

```shell
cp .env.example .env
```

### 4. Initialize the database

```shell
npx prisma generate
npx prisma migrate deploy
```

### 5. Run the dev server

```shell
npm run dev
```

### 6. Open the app in your browser

Visit [http://localhost:8000](http://localhost:3000) in your browser.

## Contributing

LeLa is an open-source project and we welcome contributions from the community.

If you'd like to contribute, please fork the repository and make changes as you'd like. Pull requests are warmly welcome.
