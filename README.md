# pingpong
Sử dụng ping để xác định trọng số (matric) cho từng node trong mạng với node management. Trọng số được so sánh giữa các lần collect khác nhau. Nếu trọng số giữa các collect biến thiên quá lớn => Network của node đang có vấn đề.
* Mô hình: pingpong agent --log matric---> Elastic Centrall --Visualation--> Network Admin (đánh giá)
* Nó mô tả lại mạng NetNORAD xác định lỗi về mặt network. Chi tiết xem: https://code.facebook.com/posts/1534350660228025/netnorad-troubleshooting-networks-via-end-to-end-probing/ 
