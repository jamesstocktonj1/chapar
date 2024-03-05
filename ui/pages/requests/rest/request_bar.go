package rest

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/mirzakhany/chapar/ui/widgets"
)

func (r *Container) requestBar(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	borderColor := widgets.Gray400
	if gtx.Source.Focused(r.address) {
		borderColor = theme.Palette.ContrastFg
	}

	border := widget.Border{
		Color:        borderColor,
		Width:        unit.Dp(1),
		CornerRadius: unit.Dp(4),
	}

	for {
		event, ok := r.address.Update(gtx)
		if !ok {
			break
		}
		if _, ok := event.(widget.ChangeEvent); ok {
			if !r.updateAddress {
				r.addressChanged()
			} else {
				r.updateAddress = false
			}
		}
	}

	// TODO fix the layout height

	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Middle,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min.Y = gtx.Dp(20)
			return layout.Inset{Right: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis:      layout.Horizontal,
						Alignment: layout.Middle,
					}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							gtx.Constraints.Min.Y = gtx.Dp(20)
							r.methodDropDown.TextSize = unit.Sp(16)
							return r.methodDropDown.Layout(gtx, theme)
						}),
						widgets.DrawLineFlex(widgets.Gray300, unit.Dp(20), unit.Dp(1)),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{Left: unit.Dp(10), Right: unit.Dp(5)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								r.addressMutex.Lock()
								defer r.addressMutex.Unlock()
								gtx.Constraints.Min.Y = gtx.Dp(20)
								editor := material.Editor(theme, r.address, "https://example.com")
								editor.TextSize = unit.Sp(16)
								return editor.Layout(gtx)
							})
						}),
					)
				})
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if r.sendClickable.Clicked(gtx) {
				go r.Submit()
			}

			gtx.Constraints.Min.X = gtx.Dp(80)
			return r.sendButton.Layout(gtx)
		}),
	)

}
