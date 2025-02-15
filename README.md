# FULLSTACK_ASSIGNMENT

### วิธีการ Migrate Data

อ้างอิงจาก <https://www.freecodecamp.org/news/database-migration-golang-migrate/>

1. ติดตั้งเครื่องมือที่ใช้ในการ migrate ผ่าน terminal ตามลิงก์ แยกตามระบบปฏิบัติการ (Windows และ macOS)
2. สร้างไฟล์ migration ในโฟลเดอร์ `db/migration/<data-base-type>` โดยใช้คำสั่ง:

   ```bash
   make migrate-create dir=[path] name=[name]
   ```

   คำสั่งนี้จะสร้างไฟล์ migration สำหรับ `up` และ `down` อย่างละหนึ่งไฟล์

3. เขียนสคริปต์สำหรับไฟล์ `up/down` ตามโครงสร้างที่ต้องการ
4. รันคำสั่ง migration:

   **migration_up:**

   ```bash
   make migrate-up path=[path] database=[database]
   ```

   **migration_down:**

   `step` คือจำนวนไฟล์ที่ต้องการ rollback (หากต้องการ rollback กลับไปก่อนหน้า 1 ไฟล์ ให้ใส่ `1`)

   ```bash
    make migrate-down path=[path] database=[database] step=[step]
   ```

5. หากเกิดข้อผิดพลาด เช่น `Dirty database version 20231006035122. Fix and force version.`
   อ้างอิงจากลิงก์ <https://stackoverflow.com/questions/59616263/dirty-database-version-error-when-using-golang-migrate>

   **แก้ไขโดยใช้ force version:**

   ```bash
    make migrate-force path=[path] database=[database] version=[version]
   ```

   จากนั้นให้รันคำสั่ง `up` เพื่อ migrate ข้อมูลอีกครั้ง:

   ```bash
    make migrate-up path=[path] database=[database]
   ```

