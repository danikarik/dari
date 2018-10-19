package parser

// Counter is described as stat collector.
type Counter int

// FieldStat contains field values.
type FieldStat struct {
	label  string
	oldval interface{}
	newval interface{}
	err    error
}

// FieldStats is a list of field stats.
type FieldStats []*FieldStat

// Stat contains information about counter and changes
// for particular type of record.
type Stat struct {
	counter    Counter
	fieldStats FieldStats
}

// Journal contains parsing stats such as inserted, updated, deleted counts.
type Journal struct {
	Total    int
	failed   *Stat
	inserted *Stat
	updated  *Stat
	deleted  *Stat
}

// NewFieldStat creates new instance of field stat.
func NewFieldStat(lbl string, old, new interface{}, err error) *FieldStat {
	return &FieldStat{
		label:  lbl,
		oldval: old,
		newval: new,
		err:    err,
	}
}

// NewStat creates new instance of stat.
func NewStat() *Stat {
	return &Stat{
		counter:    0,
		fieldStats: make(FieldStats, 0),
	}
}

// NewJournal created new instance of journal.
func NewJournal() *Journal {
	return &Journal{
		Total:    0,
		failed:   NewStat(),
		inserted: NewStat(),
		updated:  NewStat(),
		deleted:  NewStat(),
	}
}

// IncrementFailed ups failed counter
func (j *Journal) IncrementFailed(fs *FieldStat) {
	j.failed.counter++
	j.failed.fieldStats = append(j.failed.fieldStats, fs)
}

// IncrementInserted ups inserted counter
func (j *Journal) IncrementInserted(fs *FieldStat) {
	j.inserted.counter++
	j.inserted.fieldStats = append(j.inserted.fieldStats, fs)
}

// IncrementUpdated ups updated counter
func (j *Journal) IncrementUpdated(fs *FieldStat) {
	j.updated.counter++
	j.updated.fieldStats = append(j.updated.fieldStats, fs)
}

// IncrementDeleted ups deleted counter
func (j *Journal) IncrementDeleted(fs *FieldStat) {
	j.deleted.counter++
	j.deleted.fieldStats = append(j.deleted.fieldStats, fs)
}

// Failed returns counter of failed records.
func (j *Journal) Failed() Counter {
	return j.failed.counter
}

// Inserted returns counter of inserted records.
func (j *Journal) Inserted() Counter {
	return j.inserted.counter
}

// Updated returns counter of updated records.
func (j *Journal) Updated() Counter {
	return j.updated.counter
}

// Deleted returns counter of deleted records.
func (j *Journal) Deleted() Counter {
	return j.deleted.counter
}
