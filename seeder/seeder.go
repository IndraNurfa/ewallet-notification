package seeder

import (
	"ewallet-notification/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SeedTemplates(db *gorm.DB) {
	template := []models.NotificationTemplate{
		{
			TemplateName: "register",
			Subject:      "Selamat Datang di Aplikasi E-Wallet!",
			Body:         "<!DOCTYPE html><html lang='id'><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width,initial-scale=1'><title>Registrasi Berhasil</title><style>body{font-family:Arial,sans-serif;color:#333;background-color:#f4f4f4;margin:0;padding:20px}.email-container{max-width:600px;margin:0 auto;background-color:#fff;padding:20px;border-radius:8px;box-shadow:0 4px 8px rgba(0,0,0,.1)}.header{font-size:24px;font-weight:700;color:#4caf50;text-align:center;margin-bottom:20px}.content{font-size:16px;line-height:1.5;color:#666}.footer{font-size:14px;color:#999;text-align:center;margin-top:20px}</style></head><body><div class='email-container'><div class='header'>Registrasi Berhasil!</div><div class='content'><p>Halo {.full_name},</p><p>Selamat! Akun e-wallet Anda telah berhasil dibuat. Sekarang Anda dapat mulai menggunakan aplikasi kami untuk mengelola keuangan dengan mudah dan aman.</p><p>Terima kasih telah bergabung bersama kami. Jika Anda membutuhkan bantuan, jangan ragu untuk menghubungi tim dukungan kami.</p></div><div class='footer'>&copy; 2025 E-Wallet. Semua Hak Dilindungi.</div></div></body></html>",
		},
		{
			TemplateName: "purchase_success",
			Subject:      "Transaksi Berhasil",
			Body:         "<!DOCTYPE html><html><head><meta charset='UTF-8'><title>Transaksi Berhasil</title><style>body {font-family: Arial, sans-serif; background-color: #f4f4f4; margin: 0; padding: 20px; text-align: center;}.container {background-color: #fff; padding: 20px; border-radius: 8px; max-width: 600px; margin: auto;}h1 {color: #2e7d32; text-align: center;}.table {width: 80%; margin: 15px auto; border-collapse: collapse; background-color: #eef8ee; border-radius: 6px; overflow: hidden;}.table td {padding: 10px; border-bottom: 1px solid #ddd;}.table td:first-child {font-weight: bold; width: 40%;}.footer {font-size: 12px; color: #777; text-align: center; margin-top: 20px;}</style></head><body><div class='container'><h1>Transaksi Berhasil</h1><p>Halo {.full_name}, transaksi pembelian Anda telah berhasil diproses.</p><table class='table'><tr><td>Jumlah</td><td>{.amount}</td></tr><tr><td>Referensi</td><td>{.reference}</td></tr><tr><td>Deskripsi</td><td>{.description}</td></tr><tr><td>Tanggal</td><td>{.created_at}</td></tr></table><p>Terima kasih telah menggunakan layanan e-wallet kami.</p><div class='footer'>© 2025 Tim E-Wallet</div></div></body></html>",
		},
		{
			TemplateName: "purchase_failed",
			Subject:      "Transaksi Gagal",
			Body:         "<!DOCTYPE html><html><head><meta charset='UTF-8'><title>Transaksi Gagal</title><style>body {font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px; text-align: center;}.container {background-color: #fff; padding: 20px; border-radius: 8px; max-width: 600px; margin: auto;}h1 {color: #c62828; text-align: center;}.table {width: 80%; margin: 15px auto; border-collapse: collapse; background-color: #fceaea; border-radius: 6px; overflow: hidden;}.table td {padding: 10px; border-bottom: 1px solid #ddd;}.table td:first-child {font-weight: bold; width: 40%;}.footer {font-size: 12px; color: #777; text-align: center; margin-top: 20px;}</style></head><body><div class='container'><h1>Transaksi Gagal</h1><p>Halo {.full_name}, maaf, transaksi pembelian Anda tidak dapat diselesaikan.</p><table class='table'><tr><td>Jumlah</td><td>{.amount}</td></tr><tr><td>Status</td><td>{.transaction_status}</td></tr><tr><td>Alasan</td><td>{.additional_info}</td></tr><tr><td>Tanggal</td><td>{.created_at}</td></tr></table><p>Silakan coba kembali atau hubungi tim dukungan kami jika masalah berlanjut.</p><div class='footer'>© 2025 Tim E-Wallet</div></div></body></html>",
		},
		{
			TemplateName: "topup_success",
			Subject:      "Top-Up Berhasil",
			Body:         "<!DOCTYPE html><html><head><meta charset='UTF-8'><title>Top-Up Berhasil</title><style>body {font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px; text-align: center;}.container {background-color: #fff; padding: 20px; border-radius: 8px; max-width: 600px; margin: auto;}h1 {color: #2e7d32; text-align: center;}.table {width: 80%; margin: 15px auto; border-collapse: collapse; background-color: #eef8ee; border-radius: 6px; overflow: hidden;}.table td {padding: 10px; border-bottom: 1px solid #ddd;}.table td:first-child {font-weight: bold; width: 40%;}.footer {font-size: 12px; color: #777; text-align: center; margin-top: 20px;}</style></head><body><div class='container'><h1>Top-Up Berhasil</h1><p>Halo {.full_name}, saldo Anda telah berhasil ditambahkan.</p><table class='table'><tr><td>Jumlah</td><td>{.amount}</td></tr><tr><td>Referensi</td><td>{.reference}</td></tr><tr><td>Tanggal</td><td>{.created_at}</td></tr></table><p>Saldo Anda kini dapat digunakan untuk transaksi pembelian atau transfer.</p><div class='footer'>© 2025 Tim E-Wallet</div></div></body></html>",
		},
		{
			TemplateName: "topup_failed",
			Subject:      "Top-Up Gagal",
			Body:         "<!DOCTYPE html><html><head><meta charset='UTF-8'><title>Top-Up Gagal</title><style>body {font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px; text-align: center;}.container {background-color: #fff; padding: 20px; border-radius: 8px; max-width: 600px; margin: auto;}h1 {color: #c62828; text-align: center;}.table {width: 80%; margin: 15px auto; border-collapse: collapse; background-color: #fceaea; border-radius: 6px; overflow: hidden;}.table td {padding: 10px; border-bottom: 1px solid #ddd;}.table td:first-child {font-weight: bold; width: 40%;}.footer {font-size: 12px; color: #777; text-align: center; margin-top: 20px;}</style></head><body><div class='container'><h1>Top-Up Gagal</h1><p>Halo {.full_name}, permintaan top-up Anda tidak dapat diproses.</p><table class='table'><tr><td>Jumlah</td><td>{.amount}</td></tr><tr><td>Status</td><td>{.transaction_status}</td></tr><tr><td>Alasan</td><td>{.additional_info}</td></tr><tr><td>Tanggal</td><td>{.created_at}</td></tr></table><p>Silakan coba kembali atau hubungi layanan pelanggan kami untuk bantuan.</p><div class='footer'>© 2025 Tim E-Wallet</div></div></body></html>",
		},
		{
			TemplateName: "refund",
			Subject:      "Saldo Anda Telah Dikembalikan",
			Body:         "<!DOCTYPE html><html><head><meta charset='UTF-8'><title>Pengembalian Dana</title><style>body {font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px; text-align: center;}.container {background-color: #fff; padding: 20px; border-radius: 8px; max-width: 600px; margin: auto;}h1 {color: #0277bd; text-align: center;}.table {width: 80%; margin: 15px auto; border-collapse: collapse; background-color: #e7f3fc; border-radius: 6px; overflow: hidden;}.table td {padding: 10px; border-bottom: 1px solid #ddd;}.table td:first-child {font-weight: bold; width: 40%;}.footer {font-size: 12px; color: #777; text-align: center; margin-top: 20px;}</style></head><body><div class='container'><h1>Pengembalian Dana</h1><p>Halo {.full_name}, dana Anda telah berhasil dikembalikan.</p><table class='table'><tr><td>Jumlah</td><td>{.amount}</td></tr><tr><td>Referensi Asal</td><td>{.reference}</td></tr><tr><td>Deskripsi</td><td>{.description}</td></tr><tr><td>Tanggal</td><td>{.updated_at}</td></tr></table><p>Jumlah pengembalian akan muncul di saldo Anda dalam waktu singkat.</p><div class='footer'>© 2025 Tim E-Wallet</div></div></body></html>",
		},
		{
			TemplateName: "purchase_reversed",
			Subject:      "Transaksi Anda Telah Dibatalkan",
			Body:         "<!DOCTYPE html><html><head><meta charset='UTF-8'><title>Transaksi Dibatalkan</title><style>body {font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px; text-align: center;}.container {background-color: #fff; padding: 20px; border-radius: 8px; max-width: 600px; margin: auto;}h1 {color: #f9a825; text-align: center;}.table {width: 80%; margin: 15px auto; border-collapse: collapse; background-color: #fff9e6; border-radius: 6px; overflow: hidden;}.table td {padding: 10px; border-bottom: 1px solid #ddd;}.table td:first-child {font-weight: bold; width: 40%;}.footer {font-size: 12px; color: #777; text-align: center; margin-top: 20px;}</style></head><body><div class='container'><h1>Transaksi Dibatalkan</h1><p>Halo {.full_name}, transaksi pembelian Anda telah dibatalkan karena suatu kendala.</p><table class='table'><tr><td>Jumlah</td><td>{.amount}</td></tr><tr><td>Referensi</td><td>{.reference}</td></tr><tr><td>Alasan</td><td>{.additional_info}</td></tr><tr><td>Tanggal Pembatalan</td><td>{.updated_at}</td></tr></table><p>Jumlah tersebut telah dikembalikan ke saldo akun Anda.</p><div class='footer'>© 2025 Tim E-Wallet</div></div></body></html>",
		},
	}

	for _, tmpl := range template {
		var existing models.NotificationTemplate
		if err := db.Where("template_name = ?", tmpl.TemplateName).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&tmpl).Error; err != nil {
				logrus.Errorf("failed to seed %s: %v", tmpl.TemplateName, err)
			} else {
				logrus.Infof("seeded template: %s", tmpl.TemplateName)
			}
		} else {
			logrus.Infof("template already exists: %s", tmpl.TemplateName)
		}
	}
}
