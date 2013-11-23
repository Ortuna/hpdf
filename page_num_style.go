package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
*/
import "C"

type PageNumStyle int

const (
	PAGE_NUM_STYLE_DECIMAL       PageNumStyle = C.HPDF_PAGE_NUM_STYLE_DECIMAL
	PAGE_NUM_STYLE_UPPER_ROMAN   PageNumStyle = C.HPDF_PAGE_NUM_STYLE_UPPER_ROMAN
	PAGE_NUM_STYLE_LOWER_ROMAN   PageNumStyle = C.HPDF_PAGE_NUM_STYLE_LOWER_ROMAN
	PAGE_NUM_STYLE_UPPER_LETTERS PageNumStyle = C.HPDF_PAGE_NUM_STYLE_UPPER_LETTERS
	PAGE_NUM_STYLE_LOWER_LETTERS PageNumStyle = C.HPDF_PAGE_NUM_STYLE_LOWER_LETTERS
)
