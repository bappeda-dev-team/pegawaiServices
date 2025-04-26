package controller

import (
	"net/http"
	"pegawaiServices/model/web"
	"pegawaiServices/service"

	"github.com/labstack/echo/v4"
)

type PegawaiControllerImpl struct {
	pegawaiService service.PegawaiService
}

func NewPegawaiControllerImpl(pegawaiService service.PegawaiService) *PegawaiControllerImpl {
	return &PegawaiControllerImpl{pegawaiService: pegawaiService}
}

// Create godoc
// @Summary Create Pegawai
// @Description Create a new pegawai
// @Tags Pegawai
// @Accept json
// @Produce json
// @Param pegawai body web.PegawaiCreateRequest true "Create Pegawai"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /pegawai/create [post]
func (controller *PegawaiControllerImpl) Create(c echo.Context) error {
	PegawaiCreateRequest := web.PegawaiCreateRequest{}
	if err := c.Bind(&PegawaiCreateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	pegawaiResponse, err := controller.pegawaiService.Create(c.Request().Context(), PegawaiCreateRequest)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Create Employee",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Create Employee",
		Data:   pegawaiResponse,
	})
}

// Update godoc
// @Summary Update Pegawai
// @Description Update an existing pegawai
// @Tags Pegawai
// @Accept json
// @Produce json
// @Param id path string true "Pegawai ID"
// @Param pegawai body web.PegawaiUpdateRequest true "Update Pegawai"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /pegawai/update/{id} [put]
func (controller *PegawaiControllerImpl) Update(c echo.Context) error {
	PegawaiUpdateRequest := web.PegawaiUpdateRequest{}
	if err := c.Bind(&PegawaiUpdateRequest); err != nil {
		return err
	}

	id := c.Param("id")

	PegawaiUpdateRequest.Id = id

	pegawaiResponse, err := controller.pegawaiService.Update(c.Request().Context(), PegawaiUpdateRequest)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Update Employee",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Update Employee",
		Data:   pegawaiResponse,
	})
}

// Delete godoc
// @Summary Delete Pegawai
// @Description Delete an existing pegawai
// @Tags Pegawai
// @Accept json
// @Produce json
// @Param id path string true "Pegawai ID"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /pegawai/delete/{id} [delete]
func (controller *PegawaiControllerImpl) Delete(c echo.Context) error {
	id := c.Param("id")

	err := controller.pegawaiService.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Delete Employee",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Delete Employee",
		Data:   nil,
	})
}

// FindById godoc
// @Summary FindById pegawai
// @Description FindById an existing pegawai
// @Tags Pegawai
// @Accept json
// @Produce json
// @Param id path string true "Pegawai ID"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /pegawai/detail/{id} [get]
func (controller *PegawaiControllerImpl) FindById(c echo.Context) error {
	id := c.Param("id")

	pegawaiResponse, err := controller.pegawaiService.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Find Employee",
				Data:   err.Error(),
			},
		)
	}
	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find Employee",
		Data:   pegawaiResponse,
	})
}

// FindAll godoc
// @Summary FindAll pegawai
// @Description FindAll an existing pegawai
// @Tags Pegawai
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Param kode_opd query string false "Kode OPD"
// @Router /pegawai/findall [get]
func (controller *PegawaiControllerImpl) FindAll(c echo.Context) error {
	kode_opd := c.QueryParam("kode_opd")

	pegawaiResponse, err := controller.pegawaiService.FindAll(c.Request().Context(), kode_opd)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Find All Employee",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find All Employee",
		Data:   pegawaiResponse,
	})
}

// @Summary FindByNip pegawai
// @Description FindByNip an existing pegawai
// @Tags Pegawai
// @Accept json
// @Produce json
// @Param nip path string true "Employee NIP"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /pegawai/findbynip/{nip} [get]
func (controller *PegawaiControllerImpl) FindByNip(c echo.Context) error {
	pegawaiResponse, err := controller.pegawaiService.FindByNip(c.Request().Context(), c.Param("nip"))
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Find Employee by NIP",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find Employee by NIP",
		Data:   pegawaiResponse,
	})
}
