# ใช้ Node.js image
FROM node:18 AS build

# กำหนด working directory
WORKDIR /app

# คัดลอกไฟล์ dependencies
COPY package.json package-lock.json ./

# ติดตั้ง dependencies
RUN npm install

# คัดลอกโค้ดทั้งหมดเข้า container
COPY . .

# สร้างไฟล์ bundle
RUN npm run build

# ใช้ Nginx เป็นเว็บเซิร์ฟเวอร์
FROM nginx:stable-alpine

# คัดลอกไฟล์ build ไปยัง Nginx
COPY --from=build /app/dist /usr/share/nginx/html

# คัดลอกไฟล์ config เพื่อรองรับ React Router
COPY nginx.conf /etc/nginx/conf.d/default.conf

# เปิดพอร์ต
EXPOSE 80

# รัน Nginx
CMD ["nginx", "-g", "daemon off;"]
