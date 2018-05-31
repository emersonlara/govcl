// 在这里写你的事件

package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"sort"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
	"gitee.com/ying32/govcl/vcl/types/colors"
	"gitee.com/ying32/govcl/vcl/types/keys"
)

var (
	isMouseDown    bool
	mouseDownPos   types.TPoint
	saveFormSize   types.TRect
	canAcceptTypes = []string{".jpg", ".gif", ".png", ".bmp", ".ico", ".psd"}
	isAutoCenter   bool // 首次加载图片后，如果不改变ImageViewer的位置，则窗口调整大小时居中显示
	imgMousePos    types.TPoint
	orgTitle       string // 保存
	//imgBuffer      *vcl.TPicture
	ImgThumb      *vcl.TBitmap
	curDirImages  []string // 以当前打开的文件检索当前目标下的
	curImageIndex int      // 当前浏览的图片索引
)

const (
	zoomSize = 50
)

func (f *TMainForm) OnMainFormCreate(sender vcl.IObject) {
	f.SetAllowDropFiles(true)

	ImgThumb = vcl.NewBitmap()
	//f.setPBPreViewPos()
	// 调试下，先显示
	//f.PBPreview.Show()

	f.ImgViewer.SetStretch(true)
	f.ImgViewer.SetAutoSize(false)
	f.ImgIcon.Picture().Assign(vcl.Application.Icon())
	if len(os.Args) > 1 {
		if f.canAccept(os.Args[1]) {
			f.loadImage(os.Args[1])
		}
	}
	// 调试用
}

func (f *TMainForm) OnMainFormDestroy(sender vcl.IObject) {
	f.SetAllowDropFiles(false)
	ImgThumb.Free()
}

func (f *TMainForm) OnMainFormMouseWheel(sender vcl.IObject, shift types.TShiftState, wheelDelta, x, y int32, handled *bool) {
	if wheelDelta > 0 {
		f.zoom(1)
	} else {
		f.zoom(-1)
	}
}

func (f *TMainForm) getFileIndex(aFileName string) int {
	for i, fname := range curDirImages {
		if strings.Compare(fname, aFileName) == 0 {
			return i
		}
	}
	return -1
}

func (f *TMainForm) getCurrentImages(aFileName string) {
	// curDirImages
	// 清零
	curDirImages = make([]string, 0)
	path := aFileName[:len(aFileName)-len(filepath.Base(aFileName))]
	fd, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		return
	}
	defer fd.Close()
	for {
		files, err := fd.Readdir(100)
		for _, file := range files {
			if !file.IsDir() {
				fname := path + file.Name()
				if f.canAccept(fname) {
					curDirImages = append(curDirImages, fname)
				}
			}
		}
		if err == io.EOF {
			break
		}
		if len(files) == 0 {
			break
		}
	}
	// 排序下
	sort.Strings(curDirImages)
}

func (f *TMainForm) getRatio() int {
	return int(float64(f.ImgViewer.Width()) / float64(f.ImgViewer.Picture().Width()) * 100)
}

func (f *TMainForm) zoom(val int) {
	var newW, newH, newX, newY int32
	img := f.ImgViewer

	if val > 0 {
		if f.getRatio() >= 320 {
			return
		}
	} else if val < 0 {
		if f.getRatio() <= 5 {
			return
		}
	} else if val == 0 {

		newW = img.Picture().Width()
		newH = img.Picture().Height()
		ratio := 1.0
		if img.Picture().Width() > f.Width()-20 {
			newW = f.Width() - 20
			ratio = float64(newW) / float64(img.Picture().Width())
			newH = int32(float64(img.Picture().Height()) * ratio)
		} else if img.Picture().Height() > f.Height()-f.PnlTitleBar.Height()-20 {
			newH = f.Height() - f.PnlTitleBar.Height() - 20
			ratio = float64(newH) / float64(img.Picture().Height())
			newW = int32(float64(img.Picture().Width()) * ratio)
		}

		f.ImgViewer.SetBounds((f.Width()-newW)/2, (f.Height()-f.PnlTitleBar.Height()-newH)/2, newW, newH)
		// 这里调整
		//f.updateTitle()
		return
	}
	scale := 1.1
	mPos := f.ScreenToClient(vcl.Mouse.CursorPos())

	r := types.TRect{img.Left(), img.Top(), img.Left() + img.Width(), img.Top() + img.Height()}
	if r.PtInRect(mPos) {
		if val > 0 {
			newW = int32(float64(img.Width()) * scale)
			newH = int32(float64(img.Height()) * scale)
			newX = img.Left() - int32(float64(mPos.X-img.Left())*(scale-1.0))
			newY = img.Top() - int32(float64(mPos.Y-img.Top())*(scale-1.0))
		} else {
			newW = int32(float64(img.Width()) / scale)
			newH = int32(float64(img.Height()) / scale)
			newX = img.Left() + int32(float64(mPos.X-img.Left())*(1.0-1.0/scale))
			newY = img.Top() + int32(float64(mPos.Y-img.Top())*(1.0-1.0/scale))
		}
	} else {
		if val > 0 {
			newW = int32(float64(img.Width()) * scale)
			newH = int32(float64(img.Height()) * scale)
		} else {
			newW = int32(float64(img.Width()) / scale)
			newH = int32(float64(img.Height()) / scale)
		}
		newX = img.Left()
		newY = img.Top()
	}
	img.SetBounds(newX, newY, newW, newH)
	f.updateTitle()
}

func (f *TMainForm) loadImage(aFileName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
	if aFileName == "" {
		return
	}

	if strings.ToLower(path.Ext(aFileName)) == ".psd" {
		if err := PsdToBitmap(aFileName, f.ImgViewer.Picture().Bitmap()); err != nil {
			fmt.Println("err:", err)
			return
		}
	} else {
		f.ImgViewer.Picture().LoadFromFile(aFileName)
	}

	// 首次加载后居中
	isAutoCenter = true
	// 首次设置为true
	f.zoom(0.0)

	imgW := f.ImgViewer.Picture().Width()
	imgH := f.ImgViewer.Picture().Height()
	//f.ImgViewer.SetBounds((f.Width()-imgW)/2, (f.Height()-imgH)/2, imgW, imgH)

	if ImgThumb.IsValid() {
		ImgThumb.Free()
	}

	thumbW := f.PBPreview.Width()
	thumbH := int32(float64(thumbW) / float64(imgW) * float64(imgH))
	ImgThumb = vcl.NewBitmap()
	ImgThumb.SetSize(thumbW, thumbH)
	ImgThumb.Canvas().StretchDraw(types.TRect{0, 0, thumbW, thumbH}, f.ImgViewer.Picture().Graphic())

	// libvcl下可以动的gif
	if !rtl.LcLLoaded() && f.ImgViewer.Picture().Graphic().ClassName() == "TGIFImage" {
		vcl.GIFImageFromObj(f.ImgViewer.Picture().Graphic()).SetAnimate(true)
	}

	// 获取本目录下的文件名
	f.getCurrentImages(aFileName)
	curImageIndex = f.getFileIndex(aFileName)
	stat, _ := os.Stat(aFileName)
	orgTitle = fmt.Sprintf("%s (%dKB, %dx%d像素) - 第%d/%d张 ", filepath.Base(aFileName), stat.Size()/1024, imgW, imgH, curImageIndex+1, len(curDirImages))
	f.updateTitle()

}

func (f *TMainForm) OnMainFormResize(sender vcl.IObject) {
	// 首次加载后居中
	if isAutoCenter {
		fmt.Println("isAutoCenter:", isAutoCenter)
		f.zoom(0)
	}
	f.setPBPreViewPos()
}

func (f *TMainForm) OnMainFormKeyDown(sender vcl.IObject, key *types.Char, shift types.TShiftState) {

	/*
	  vkLeft             = 0x25;  //  37
	  vkUp               = 0x26;  //  38
	  vkRight            = 0x27;  //  39
	  vkDown             = 0x28;  //  40
	*/
	if len(curDirImages) == 0 {
		return
	}

	switch *key {
	// Left
	case keys.VkLeft:
		curImageIndex--
		if curImageIndex < 0 {
			curImageIndex = len(curDirImages) - 1
		}
		f.loadImage(curDirImages[curImageIndex])

	// Right
	case keys.VkRight:
		curImageIndex++
		if curImageIndex >= len(curDirImages) {
			curImageIndex = 0
		}
		f.loadImage(curDirImages[curImageIndex])
	}
}

func (f *TMainForm) setPBPreViewPos() {
	if f.PBPreview.Visible() {
		f.PBPreview.SetLeft(f.Width() - f.PBPreview.Width() - 20)
		f.PBPreview.SetTop(f.Height() - f.PBPreview.Height() - 50)
	}
}

func (f *TMainForm) updateTitle() {
	title := orgTitle
	ratio := f.getRatio()
	if ratio != 100 {
		title += fmt.Sprintf(" - %d%%", f.getRatio())
	}
	title += " - ying32图片浏览器"

	f.SetCaption(title)
	f.LblCaption.SetCaption(title)
}

func (f *TMainForm) canAccept(aFileName string) bool {
	ext := strings.ToLower(filepath.Ext(aFileName))
	for _, s := range canAcceptTypes {
		if ext == s {
			return true
		}
	}
	return false
}

func (f *TMainForm) isMax() bool {
	return f.Width() == vcl.Screen.WorkAreaWidth() && f.Height() == vcl.Screen.WorkAreaHeight()
}

func (f *TMainForm) OnBtnMinClick(sender vcl.IObject) {
	vcl.Application.Minimize()
}

func (f *TMainForm) OnBtnCloseClick(sender vcl.IObject) {
	f.Close()
}

func (f *TMainForm) OnBtnMaxClick(sender vcl.IObject) {
	f.BtnMax.Hide()
	f.BtnRestore.Show()
	saveFormSize.Left = f.Left()
	saveFormSize.Top = f.Top()
	saveFormSize.Right = saveFormSize.Left + f.Width()
	saveFormSize.Bottom = saveFormSize.Top + f.Height()
	f.SetBounds(0, 0, vcl.Screen.WorkAreaWidth(), vcl.Screen.WorkAreaHeight())
}

func (f *TMainForm) OnBtnRestoreClick(sender vcl.IObject) {
	f.BtnRestore.Hide()
	f.BtnMax.Show()
	f.SetBounds(saveFormSize.Left, saveFormSize.Top, saveFormSize.Width(), saveFormSize.Height())
}

func (f *TMainForm) OnBtnMenuClick(sender vcl.IObject) {
	vcl.ShowMessage("别点了，点了我我也不干活！")
}

func (f *TMainForm) OnLblCaptionMouseUp(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	isMouseDown = false
}

func (f *TMainForm) OnLblCaptionMouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if button == types.MbLeft {
		// windows下可以用这种，貌似只对libvcl生效，lcl不生效的
		//win.ReleaseCapture()
		//f.Perform(win.WM_SYSCOMMAND, win.SC_MOVE+1, 0)
		isMouseDown = true
		mouseDownPos.X = x
		mouseDownPos.Y = y
	}
}

func (f *TMainForm) OnLblCaptionMouseMove(sender vcl.IObject, shift types.TShiftState, x, y int32) {
	if isMouseDown {
		if f.isMax() {
			f.BtnRestore.Click()
			// 模拟的就是麻烦啊，要各种计算
			ratio := float64(saveFormSize.Width()) / float64(vcl.Screen.WorkAreaWidth())
			x = int32(float64(x) * ratio)
			mouseDownPos.X = x
		}
		f.SetLeft(f.Left() + (x - mouseDownPos.X))
		f.SetTop(f.Top() + (y - mouseDownPos.Y))
	}
}

func (f *TMainForm) OnLblCaptionDblClick(sender vcl.IObject) {
	//if f.isMax() {
	//	f.BtnRestore.Click()
	//} else {
	//	f.BtnMax.Click()
	//}
}

func (f *TMainForm) OnImgViewerMouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if button == types.MbLeft {
		isMouseDown = true
		mouseDownPos.X = x
		mouseDownPos.Y = y
	}
}

func (f *TMainForm) OnImgViewerMouseUp(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	isMouseDown = false
}

func (f *TMainForm) OnImgViewerMouseMove(sender vcl.IObject, shift types.TShiftState, x, y int32) {
	imgMousePos.X = x
	imgMousePos.Y = y
	if isMouseDown {
		isAutoCenter = false
		f.ImgViewer.SetLeft(f.ImgViewer.Left() + (x - mouseDownPos.X))
		f.ImgViewer.SetTop(f.ImgViewer.Top() + (y - mouseDownPos.Y))
	}
}

func (f *TMainForm) OnMainFormDropFiles(sender vcl.IObject, aFileNames []string) {
	// 只第一个
	if f.canAccept(aFileNames[0]) {
		f.loadImage(aFileNames[0])
	}
}

func (f *TMainForm) OnPBPreviewMouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {

}

func (f *TMainForm) OnPBPreviewPaint(sender vcl.IObject) {
	if ImgThumb.IsValid() {
		pb := f.PBPreview
		canvas := pb.Canvas()
		canvas.Pen().SetColor(colors.ClBlack)
		canvas.Brush().SetStyle(types.BsClear)
		canvas.Rectangle(0, 0, pb.Width(), pb.Height())
		canvas.Draw(0, pb.Height()-ImgThumb.Height(), ImgThumb)
	}
}

func (f *TMainForm) OnPBPreviewMouseMove(sender vcl.IObject, shift types.TShiftState, x, y int32) {

}

func (f *TMainForm) OnPBPreviewMouseUp(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {

}