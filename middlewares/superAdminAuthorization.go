package middlewares

import (
	"errors"

	"net/http"

	"backend_server/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

// SuperAdminMiddleware memastikan hanya Super Admin yang boleh mengakses route tertentu
func SuperAdminMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		// Ambil koneksi DB dari context

		tx, ok := c.Value("tx").(*pop.Connection)
		if !ok {
			return c.Error(http.StatusInternalServerError, errors.New("database connection not found"))
		}

		// Ambil ID admin dari context (hasil validasi JWT)
		adminID, ok := c.Value("admin_id").(string)
		if !ok || adminID == "" {
			return c.Error(http.StatusUnauthorized, errors.New("unauthorized - no admin id in context"))
		}

		// Ambil data admin dari database
		admin := &models.Admin{}
		if err := tx.Find(admin, adminID); err != nil {
			return c.Error(http.StatusUnauthorized, errors.New("admin not found"))
		}

		// Ambil data role admin
		role := &models.Role{}
		if err := tx.Find(role, admin.RoleID); err != nil {
			return c.Error(http.StatusUnauthorized, errors.New("role not found"))
		}

		// Cek apakah role-nya super admin
		if role.Name != "super_admin" {
			return c.Error(http.StatusForbidden, errors.New("forbidden: only super admin allowed"))
		}

		// Jika semua aman, lanjutkan ke handler berikutnya
		return next(c)
	}
}
