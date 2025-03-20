# フロントエンド用のDockerfile
FROM node:20-alpine

WORKDIR /app

# package.json だけコピーして、依存関係を先にインストール
COPY package.json package-lock.json* ./
RUN npm install

# `vite` がインストールされているか確認
RUN npx vite --version

# 残りのファイルをコピー
COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev"]
