package mongo

type Update struct {
}

func (u *Update) Set(fieldName string, value interface{}) *Update   {}
func (u *Update) UnSet(fieldName string, value interface{}) *Update {}

func (u *Update) Increment(fieldName string, value interface{}) *Update {}

func (u *Update) Pull(fieldName string, value interface{}) *Update {}

func (u *Update) Push(fieldName string, value interface{}) *Update {}

func (u *Update) Pop(fieldName string, value interface{}) *Update {}

func (u *Update) Max(fieldName string, value interface{}) *Update {}
func (u *Update) Min(fieldName string, value interface{}) *Update {}
func (u *Update) Mul(fieldName string, value interface{}) *Update {}

func (u *Update) Rename(fieldName string, value interface{}) *Update {}
