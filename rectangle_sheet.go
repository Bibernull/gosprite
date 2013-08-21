package gosprite
import(
	"time"
	"fmt"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

type RectangleSheet struct{
    width int
    height int
    data TwoDimensional
}

func (r *RectangleSheet) init(width, height int){
	r.width = width
	r.height = height
	r.data = TwoDimensional{}
	r.data.init(width, height)
}


func (r *RectangleSheet) canAddRectangleAt(x, y, width, height int) (requiredHorisontalCount, requiredVerticalCount, leftOverWidth, leftOverHeight int, err error){
	requiredHorisontalCount = 0;
	requiredVerticalCount = 0;
	leftOverWidth = 0;
	leftOverHeight = 0;

	foundWidth := 0;
	foundHeight := 0;
	trialX := x;
	trialY := y;
	err = nil
	// Check all cells that need to be unoccupied for there to be room for the rectangle.
	//console.log("Occupied " + this.data.item(trialX, trialY).occupied + " at " + trialX + " " + trialY);

	for foundHeight < height {
	    trialX = x;
	    foundWidth = 0;

	    for foundWidth < width {
	        if r.data.item(trialX, trialY) == true {
	        	err =  MyError{
							time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
							"Cannot place Rectangle",
						}
	            return ;
	        }

	        foundWidth += r.data.colWidth(trialX);
	        trialX++;
	    }

	    foundHeight += r.data.rowHeight(trialY);
	    trialY++;
	}

	// Visited all cells that we'll need to place the rectangle,
	// and none were occupied. So the space is available here.

	requiredHorisontalCount = trialX - x;
	requiredVerticalCount = trialY - y;

	leftOverWidth = (foundWidth - width);
	leftOverHeight = (foundHeight - height);

	return

	};

func (r *RectangleSheet) addRectangle(height, width int) (offsetX, offsetY int) {

	requiredHeight := height;
	requiredWidth := width;

	x := 0;
	y := 0;
	offsetX = 0;
	offsetY = 0;

	rowCount := r.data.rowCount;

	for {

	    // First move upwards until we find an unoccupied cell.
	    // If we're already at an unoccupied cell, no need to do anything.
	    // Important to clear all occupied cells to get
	    // the lowest free height deficit. This must be taken from the top of the highest
	    // occupied cell.

	    for (y < rowCount) && (r.data.item(x, y) == true) {
	        offsetY += r.data.rowHeight(y)
	        y += 1;
	    }

	    // If we found an unoccupied cell, than see if we can place a rectangle there.
	    // If not, than y popped out of the top of the canvas.

	    if ((y < rowCount) && (r.freeHeightDeficit(r.height, offsetY, requiredHeight) <= 0)) {

          requiredHorisontalCount, requiredVerticalCount, leftOverWidth, leftOverHeight, err := r.canAddRectangleAt(x, y, requiredWidth, requiredHeight);

	        if err == nil {

	            r.placeRectangle(x, y, requiredWidth,requiredHeight,
	                requiredHorisontalCount, requiredVerticalCount,
	                leftOverWidth, leftOverHeight);


	            return;
	        }
	        // Go to the next cell
	        offsetY += r.data.rowHeight(y)
	        y += 1;
	    }

	    // If we've come so close to the top of the canvas that there is no space for the
	    // rectangle, go to the next column. r automatically also checks whether we've popped out of the top
	    // of the canvas (in that case, _canvasHeight == offsetY).

	    var freeHeightDeficit = r.freeHeightDeficit(r.height, offsetY, requiredHeight);
	    if freeHeightDeficit > 0 {
	        offsetY = 0;
	        y = 0;

	        offsetX += r.data.colWidth(x);

	        x += 1;

	    }

	    // If we've come so close to the right edge of the canvas that there is no space for
	    // the rectangle, return false now.
	    if (r.width - offsetX) < requiredWidth {
	        return -1, -1;
	    }
	}

	return
}

func (r *RectangleSheet) freeHeightDeficit(canvasHeight, offsetY, requiredHeight int) int {
    spaceLeftVertically := canvasHeight - offsetY
    freeHeightDeficit := requiredHeight - spaceLeftVertically

    return freeHeightDeficit
}

func (r *RectangleSheet) placeRectangle(x, y, height, width, requiredHorisontalCount, requiredVerticalCount, leftOverWidth, leftOverHeight int) {
    if leftOverWidth > 0 {
        xFarRightColumn := x + requiredHorisontalCount - 1
        r.data.insertColumn(xFarRightColumn, leftOverWidth)
    }

    if leftOverHeight > 0 {
        yFarBottomColumn := y + requiredVerticalCount - 1
        r.data.insertRow(yFarBottomColumn, leftOverHeight)
    }

    for i := x + requiredHorisontalCount - 1; i >= x; i-- {
        for j := y + requiredVerticalCount - 1; j >= y; j-- {
            r.data.setItem(i, j, true)
        }
    }

}

func (r *RectangleSheet) calculate(images []*Image) {

	totalWidth := 0
    maxHeight := images[0].height + (images[0].height/2)

    for _, v := range images {
        totalWidth += v.width
    }

    r.init(totalWidth, maxHeight)

    success_count := 0;

    for _, image := range images{
        offsetX, offsetY := r.addRectangle(image.height, image.width);
        if offsetX >= 0 && offsetY >= 0 {

            image.x = offsetX;
            image.y = offsetY;

            success_count++;
        }

    }

    last_col_occupied := false;
    for y := 0; y < r.data.rowCount; y++ {
        if r.data.item(r.data.colCount - 1, y) == true{
            last_col_occupied = true;
        }
    }

    if last_col_occupied == false{
        r.width -= r.data.colWidth(r.data.colCount - 1);
    }

    return
};