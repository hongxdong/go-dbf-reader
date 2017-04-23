package dbf

import (
  "os"
  "testing"
  "fmt"
)

func TestRead(t *testing.T) {
  f, err := os.Open("./show2003.dbf")
  // f, err := os.Open("./sjshq.dbf")
  if err != nil {
    t.Fatal(err)
  }

  dbr, err := NewReader(f)
  if err != nil {
    t.Fatal(err)
  }
  fmt.Printf("Mod date: %d-%d-%d\n", dbr.year, dbr.month, dbr.day)
  fmt.Printf("RecordLen:%d\n", dbr.recordlen)
  fmt.Printf("RecordNum:%d\n", dbr.Length)

  field_names := dbr.FieldNames()
  for i, name := range field_names {
    fmt.Print(name, ":")
    fmt.Printf("%s", string(dbr.fields[i].Type))
    fmt.Print(":", dbr.fields[i].Len, ":", dbr.fields[i].DecimalPlaces, ",")
  }
  fmt.Println("--->")


  //rec, e := dbr.Read(0)
  //fmt.Println(rec, e, i)


  for i := uint32(0); i < uint32(dbr.Length) ; i++ {
    record, err := dbr.Read(i)
    if err != nil && err.Error() == "EOF" {
      break
    }
    fmt.Print(i, ": ", err)
    fmt.Println(record)
  }

}
