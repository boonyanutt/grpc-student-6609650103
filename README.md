# grpc-student-6609650103

## สิ่งที่เพิ่มเข้ามา (Assignments)
- **Task 1 (ListStudents)**: เพิ่มเมธอดสำหรับดึงรายชื่อนักศึกษาทั้งหมดในระบบ
- **Task 2 (Phone Field)**: เพิ่มฟิลด์เบอร์โทรศัพท์ (phone) ในข้อมูลนักศึกษา

## โครงสร้างโปรเจกต์
- `proto/`: ไฟล์คำจำกัดความของ Service (`.proto`)
- `server/`: ส่วนการทำงานของ Server
- `client/`: ส่วนการเรียกใช้งานของ Client เพื่อทดสอบระบบ
- `studentpb/`: โค้ดที่ Generate ออกมาจากไฟล์ Proto

## วิธีการรันระบบ
1. เปิด Terminal แรกเพื่อรัน Server: `go run server/server.go`
2. เปิด Terminal ที่สองเพื่อรัน Client: `go run client/client.go`