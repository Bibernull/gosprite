package gosprite
import (
    "fmt"
)

type Dimension struct{
	size int
	index int
}

type TwoDimensional struct{
    columns []Dimension
	rows []Dimension
    data [][]bool

    rowCount int
    colCount int
}

func (t *TwoDimensional) init(width, height int) {
	t.columns = make([]Dimension, width)
	t.rows = make([]Dimension, height)
	t.data = make([][]bool,width)

	t.rows[0] = Dimension{size: height, index: 0}
	t.columns[0] = Dimension{size: width, index: 0}

	for key, _ := range t.data {
		t.data[key] = make([]bool, height)
	}
	t.rowCount = 1
	t.colCount = 1
}

func (t *TwoDimensional) insertRow (y, height int) {
    if y >= t.rowCount {
        fmt.Print("Row count is ")
        fmt.Print(t.rowCount)
        fmt.Print("attempted insert at ")
        fmt.Print(y)
        return
	}

    // Copy the cells with the given y to a new row after the last used row. The y of the new row equals _nbrRows.

    var physY = t.rows[y].index;

    for x := 0; x < t.colCount; x++ {
            //this.data[x][this.rowCount] = (this.data[x][physY]) ? true : false;
            t.data[x][t.rowCount] = t.data[x][physY];

    }

    if y < t.rowCount - 1{
        for i := t.rowCount; i > y + 1; i--{
            t.rows[i] = t.rows[i - 1];

           /* for(j = 0; j < this.colCount; j++){
                this.setItem(j,i, this.item(j, i - 1))
            }*/
        }
    }

    var oldH = t.rows[y].size;
    var newH = oldH - height;
    t.rows[y + 1] = Dimension{size: height, index: t.rowCount};
    t.rows[y].size = newH;
    t.rowCount++;
}

func (t *TwoDimensional) insertColumn (x, width int) {

	if x >= t.colCount {
	    fmt.Print("Column count is ")
	    fmt.Print(t.colCount)
	    fmt.Print("attempted insert at ")
	    fmt.Print(x)
	    return
	}

    physRowCount := len(t.data[0]);
    var physX = t.columns[x].index;
    t.data[t.colCount] = make([]bool, physRowCount);

    for y := 0; y < physRowCount;y++{
        t.data[t.colCount][y] = (t.data[physX][y]);
    }


    // Make room in the _columns array by shifting all items that come after the one indexed by x one position to the right.
    // If x is at the end of the array, there is no need to shift anything.
    if x < (t.colCount - 1) {
        for i := t.colCount; i > x; i--{
            t.columns[i] = t.columns[i - 1];

           }
    }

    // Set the widths of the old and new columns.
    var oldW = t.columns[x].size;
    var newW = oldW - width;

    t.columns[x + 1] = Dimension{size: width, index: t.colCount}
    t.columns[x].size = newW;

    // The logical width of the array has increased by 1.
    t.colCount++;
}

func (t *TwoDimensional) setItem (x, y int, value bool) {
    t.data[t.columns[x].index][t.rows[y].index] = value;
};

func (t TwoDimensional) item (x, y int) bool {
     return t.data[t.columns[x].index][t.rows[y].index];
};

func (t TwoDimensional) rowHeight (y int) int{
    return t.rows[y].size;
};

func (t TwoDimensional) colWidth (x int) int{
    return t.columns[x].size;
};