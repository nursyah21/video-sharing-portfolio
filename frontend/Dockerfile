
FROM node:22-alpine AS build

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn install

COPY . ./

RUN yarn run build


FROM nginx:stable-alpine AS production

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY nginx.conf /etc/nginx/conf.d

COPY --from=build /app/dist ./

ENTRYPOINT ["nginx", "-g", "daemon off;"]


