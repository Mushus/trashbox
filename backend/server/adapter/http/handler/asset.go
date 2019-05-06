package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/Mushus/trashbox/backend/server/app/asset"
	"github.com/labstack/echo/v4"
)

func (h Handler) UploadAsset(c Context) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return err
	}

	fileName := fileHeader.Filename
	contentType := fileHeader.Header.Get("Content-Type")
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	sa := asset.Asset{
		Stream:      file,
		FileName:    fileName,
		ContentType: contentType,
	}
	id, err := h.asset.Add(sa)
	if err != nil {
		return err
	}

	defer file.Close()

	return c.String(http.StatusOK, id)
}

func (h Handler) decolateAssetResponse(c Context, asset asset.Asset) error {
	resp := c.Response()
	header := resp.Header()
	// for download
	download := c.QueryParam("download")
	if download == "true" || download == "yes" {
		encodedFileName := url.QueryEscape(asset.FileName)
		hValue := fmt.Sprintf(`attachment;filename*="UTF-8''%s"`, encodedFileName)
		header.Set(echo.HeaderContentDisposition, hValue)
	}
	return nil
}

// GetAsset is a handler
func (h Handler) GetAsset(c Context) error {
	id := c.Param("id")

	a, err := h.asset.Get(id)
	if err == asset.AssetNotFound {
		return c.String(http.StatusNotFound, "asset not found")
	}
	if err != nil {
		return err
	}
	defer a.Close()

	h.decolateAssetResponse(c, a)
	return c.Stream(http.StatusOK, a.ContentType, a)
}

// GetFormatedAsset is a handler
func (h Handler) GetFormatedAsset(c Context) error {
	id := c.Param("id")
	// format := c.Param("format")

	a, err := h.asset.Get(id)
	if err == asset.AssetNotFound {
		return c.String(http.StatusNotFound, "asset not found")
	}
	if err != nil {
		return err
	}
	defer a.Close()

	h.decolateAssetResponse(c, a)
	return c.Stream(http.StatusOK, a.ContentType, a)
}
