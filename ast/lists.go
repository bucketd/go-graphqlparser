// Code generated by lab/generics
// DO NOT EDIT!
package ast

// Arguments is a linked list that contains Argument values.
type Arguments struct {
	Data Argument
	next *Arguments
	pos  int
}

// Add appends a Argument to this linked list and returns this new head.
func (as *Arguments) Add(data Argument) *Arguments {
	var pos int

	if as != nil {
		pos = as.pos + 1
	}

	return &Arguments{
		Data: data,
		next: as,
		pos:  pos,
	}
}

// ForEach applies the given map function to each item in this linked list.
func (as *Arguments) ForEach(fn func(a Argument, i int)) {
	if as == nil {
		return
	}

	iter := 0
	current := as

	for {
		fn(current.Data, iter)

		if current.next == nil {
			break
		}

		iter++
		current = current.next
	}
}

// Insert places the Argument in the position given by pos.
// The method will insert at top if pos is greater than or equal to list length.
// The method will insert at bottom if the pos is less than 0.
func (as *Arguments) Insert(a Argument, pos int) *Arguments {
	if pos >= as.Len() || as == nil {
		return as.Add(a)
	}

	if pos < 0 {
		pos = 0
	}

	mid := as
	for mid.pos != pos {
		mid = mid.next
	}

	bot := mid.next
	mid.next = nil
	as.pos -= mid.pos

	bot = bot.Add(a)
	as.Join(bot)

	return as
}

// Join attaches the tail of the receiver list "as" to the head of the otherList.
func (as *Arguments) Join(otherList *Arguments) {
	if as == nil {
		return
	}

	pos := as.Len() + otherList.Len() - 1

	last := as
	for as != nil {
		as.pos = pos
		pos--
		last = as
		as = as.next
	}

	last.next = otherList
}

// Len returns the length of this linked list.
func (as *Arguments) Len() int {
	if as == nil {
		return 0
	}
	return as.pos + 1
}

// Reverse reverses this linked list of Argument. Usually when the linked list is being
// constructed the result will be last-to-first, so we'll want to reverse it to get it in the
// "right" order.
func (as *Arguments) Reverse() *Arguments {
	current := as

	var prev *Arguments
	var pos int

	for current != nil {
		current.pos = pos
		pos++

		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

// ArgumentsFromSlice returns a Arguments list from a slice of Argument.
func ArgumentsFromSlice(sl []Argument) *Arguments {
	var list *Arguments
	for _, v := range sl {
		list = list.Add(v)
	}
	return list.Reverse()
}

// Definitions is a linked list that contains Definition values.
type Definitions struct {
	Data Definition
	next *Definitions
	pos  int
}

// Add appends a Definition to this linked list and returns this new head.
func (ds *Definitions) Add(data Definition) *Definitions {
	var pos int

	if ds != nil {
		pos = ds.pos + 1
	}

	return &Definitions{
		Data: data,
		next: ds,
		pos:  pos,
	}
}

// ForEach applies the given map function to each item in this linked list.
func (ds *Definitions) ForEach(fn func(d Definition, i int)) {
	if ds == nil {
		return
	}

	iter := 0
	current := ds

	for {
		fn(current.Data, iter)

		if current.next == nil {
			break
		}

		iter++
		current = current.next
	}
}

// Insert places the Definition in the position given by pos.
// The method will insert at top if pos is greater than or equal to list length.
// The method will insert at bottom if the pos is less than 0.
func (ds *Definitions) Insert(d Definition, pos int) *Definitions {
	if pos >= ds.Len() || ds == nil {
		return ds.Add(d)
	}

	if pos < 0 {
		pos = 0
	}

	mid := ds
	for mid.pos != pos {
		mid = mid.next
	}

	bot := mid.next
	mid.next = nil
	ds.pos -= mid.pos

	bot = bot.Add(d)
	ds.Join(bot)

	return ds
}

// Join attaches the tail of the receiver list "ds" to the head of the otherList.
func (ds *Definitions) Join(otherList *Definitions) {
	if ds == nil {
		return
	}

	pos := ds.Len() + otherList.Len() - 1

	last := ds
	for ds != nil {
		ds.pos = pos
		pos--
		last = ds
		ds = ds.next
	}

	last.next = otherList
}

// Len returns the length of this linked list.
func (ds *Definitions) Len() int {
	if ds == nil {
		return 0
	}
	return ds.pos + 1
}

// Reverse reverses this linked list of Definition. Usually when the linked list is being
// constructed the result will be last-to-first, so we'll want to reverse it to get it in the
// "right" order.
func (ds *Definitions) Reverse() *Definitions {
	current := ds

	var prev *Definitions
	var pos int

	for current != nil {
		current.pos = pos
		pos++

		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

// DefinitionsFromSlice returns a Definitions list from a slice of Definition.
func DefinitionsFromSlice(sl []Definition) *Definitions {
	var list *Definitions
	for _, v := range sl {
		list = list.Add(v)
	}
	return list.Reverse()
}

// Directives is a linked list that contains Directive values.
type Directives struct {
	Data Directive
	next *Directives
	pos  int
}

// Add appends a Directive to this linked list and returns this new head.
func (ds *Directives) Add(data Directive) *Directives {
	var pos int

	if ds != nil {
		pos = ds.pos + 1
	}

	return &Directives{
		Data: data,
		next: ds,
		pos:  pos,
	}
}

// ForEach applies the given map function to each item in this linked list.
func (ds *Directives) ForEach(fn func(d Directive, i int)) {
	if ds == nil {
		return
	}

	iter := 0
	current := ds

	for {
		fn(current.Data, iter)

		if current.next == nil {
			break
		}

		iter++
		current = current.next
	}
}

// Insert places the Directive in the position given by pos.
// The method will insert at top if pos is greater than or equal to list length.
// The method will insert at bottom if the pos is less than 0.
func (ds *Directives) Insert(d Directive, pos int) *Directives {
	if pos >= ds.Len() || ds == nil {
		return ds.Add(d)
	}

	if pos < 0 {
		pos = 0
	}

	mid := ds
	for mid.pos != pos {
		mid = mid.next
	}

	bot := mid.next
	mid.next = nil
	ds.pos -= mid.pos

	bot = bot.Add(d)
	ds.Join(bot)

	return ds
}

// Join attaches the tail of the receiver list "ds" to the head of the otherList.
func (ds *Directives) Join(otherList *Directives) {
	if ds == nil {
		return
	}

	pos := ds.Len() + otherList.Len() - 1

	last := ds
	for ds != nil {
		ds.pos = pos
		pos--
		last = ds
		ds = ds.next
	}

	last.next = otherList
}

// Len returns the length of this linked list.
func (ds *Directives) Len() int {
	if ds == nil {
		return 0
	}
	return ds.pos + 1
}

// Reverse reverses this linked list of Directive. Usually when the linked list is being
// constructed the result will be last-to-first, so we'll want to reverse it to get it in the
// "right" order.
func (ds *Directives) Reverse() *Directives {
	current := ds

	var prev *Directives
	var pos int

	for current != nil {
		current.pos = pos
		pos++

		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

// DirectivesFromSlice returns a Directives list from a slice of Directive.
func DirectivesFromSlice(sl []Directive) *Directives {
	var list *Directives
	for _, v := range sl {
		list = list.Add(v)
	}
	return list.Reverse()
}

// DirectiveLocations is a linked list that contains DirectiveLocation values.
type DirectiveLocations struct {
	Data DirectiveLocation
	next *DirectiveLocations
	pos  int
}

// Add appends a DirectiveLocation to this linked list and returns this new head.
func (dls *DirectiveLocations) Add(data DirectiveLocation) *DirectiveLocations {
	var pos int

	if dls != nil {
		pos = dls.pos + 1
	}

	return &DirectiveLocations{
		Data: data,
		next: dls,
		pos:  pos,
	}
}

// ForEach applies the given map function to each item in this linked list.
func (dls *DirectiveLocations) ForEach(fn func(dl DirectiveLocation, i int)) {
	if dls == nil {
		return
	}

	iter := 0
	current := dls

	for {
		fn(current.Data, iter)

		if current.next == nil {
			break
		}

		iter++
		current = current.next
	}
}

// Insert places the DirectiveLocation in the position given by pos.
// The method will insert at top if pos is greater than or equal to list length.
// The method will insert at bottom if the pos is less than 0.
func (dls *DirectiveLocations) Insert(dl DirectiveLocation, pos int) *DirectiveLocations {
	if pos >= dls.Len() || dls == nil {
		return dls.Add(dl)
	}

	if pos < 0 {
		pos = 0
	}

	mid := dls
	for mid.pos != pos {
		mid = mid.next
	}

	bot := mid.next
	mid.next = nil
	dls.pos -= mid.pos

	bot = bot.Add(dl)
	dls.Join(bot)

	return dls
}

// Join attaches the tail of the receiver list "dls" to the head of the otherList.
func (dls *DirectiveLocations) Join(otherList *DirectiveLocations) {
	if dls == nil {
		return
	}

	pos := dls.Len() + otherList.Len() - 1

	last := dls
	for dls != nil {
		dls.pos = pos
		pos--
		last = dls
		dls = dls.next
	}

	last.next = otherList
}

// Len returns the length of this linked list.
func (dls *DirectiveLocations) Len() int {
	if dls == nil {
		return 0
	}
	return dls.pos + 1
}

// Reverse reverses this linked list of DirectiveLocation. Usually when the linked list is being
// constructed the result will be last-to-first, so we'll want to reverse it to get it in the
// "right" order.
func (dls *DirectiveLocations) Reverse() *DirectiveLocations {
	current := dls

	var prev *DirectiveLocations
	var pos int

	for current != nil {
		current.pos = pos
		pos++

		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

// DirectiveLocationsFromSlice returns a DirectiveLocations list from a slice of DirectiveLocation.
func DirectiveLocationsFromSlice(sl []DirectiveLocation) *DirectiveLocations {
	var list *DirectiveLocations
	for _, v := range sl {
		list = list.Add(v)
	}
	return list.Reverse()
}

// EnumValueDefinitions is a linked list that contains EnumValueDefinition values.
type EnumValueDefinitions struct {
	Data EnumValueDefinition
	next *EnumValueDefinitions
	pos  int
}

// Add appends a EnumValueDefinition to this linked list and returns this new head.
func (evds *EnumValueDefinitions) Add(data EnumValueDefinition) *EnumValueDefinitions {
	var pos int

	if evds != nil {
		pos = evds.pos + 1
	}

	return &EnumValueDefinitions{
		Data: data,
		next: evds,
		pos:  pos,
	}
}

// ForEach applies the given map function to each item in this linked list.
func (evds *EnumValueDefinitions) ForEach(fn func(evd EnumValueDefinition, i int)) {
	if evds == nil {
		return
	}

	iter := 0
	current := evds

	for {
		fn(current.Data, iter)

		if current.next == nil {
			break
		}

		iter++
		current = current.next
	}
}

// Insert places the EnumValueDefinition in the position given by pos.
// The method will insert at top if pos is greater than or equal to list length.
// The method will insert at bottom if the pos is less than 0.
func (evds *EnumValueDefinitions) Insert(evd EnumValueDefinition, pos int) *EnumValueDefinitions {
	if pos >= evds.Len() || evds == nil {
		return evds.Add(evd)
	}

	if pos < 0 {
		pos = 0
	}

	mid := evds
	for mid.pos != pos {
		mid = mid.next
	}

	bot := mid.next
	mid.next = nil
	evds.pos -= mid.pos

	bot = bot.Add(evd)
	evds.Join(bot)

	return evds
}

// Join attaches the tail of the receiver list "evds" to the head of the otherList.
func (evds *EnumValueDefinitions) Join(otherList *EnumValueDefinitions) {
	if evds == nil {
		return
	}

	pos := evds.Len() + otherList.Len() - 1

	last := evds
	for evds != nil {
		evds.pos = pos
		pos--
		last = evds
		evds = evds.next
	}

	last.next = otherList
}

// Len returns the length of this linked list.
func (evds *EnumValueDefinitions) Len() int {
	if evds == nil {
		return 0
	}
	return evds.pos + 1
}

// Reverse reverses this linked list of EnumValueDefinition. Usually when the linked list is being
// constructed the result will be last-to-first, so we'll want to reverse it to get it in the
// "right" order.
func (evds *EnumValueDefinitions) Reverse() *EnumValueDefinitions {
	current := evds

	var prev *EnumValueDefinitions
	var pos int

	for current != nil {
		current.pos = pos
		pos++

		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

// EnumValueDefinitionsFromSlice returns a EnumValueDefinitions list from a slice of EnumValueDefinition.
func EnumValueDefinitionsFromSlice(sl []EnumValueDefinition) *EnumValueDefinitions {
	var list *EnumValueDefinitions
	for _, v := range sl {
		list = list.Add(v)
	}
	return list.Reverse()
}

// FieldDefinitions is a linked list that contains FieldDefinition values.
type FieldDefinitions struct {
	Data FieldDefinition
	next *FieldDefinitions
	pos  int
}

// Add appends a FieldDefinition to this linked list and returns this new head.
func (fds *FieldDefinitions) Add(data FieldDefinition) *FieldDefinitions {
	var pos int

	if fds != nil {
		pos = fds.pos + 1
	}

	return &FieldDefinitions{
		Data: data,
		next: fds,
		pos:  pos,
	}
}

// ForEach applies the given map function to each item in this linked list.
func (fds *FieldDefinitions) ForEach(fn func(fd FieldDefinition, i int)) {
	if fds == nil {
		return
	}

	iter := 0
	current := fds

	for {
		fn(current.Data, iter)

		if current.next == nil {
			break
		}

		iter++
		current = current.next
	}
}

// Insert places the FieldDefinition in the position given by pos.
// The method will insert at top if pos is greater than or equal to list length.
// The method will insert at bottom if the pos is less than 0.
func (fds *FieldDefinitions) Insert(fd FieldDefinition, pos int) *FieldDefinitions {
	if pos >= fds.Len() || fds == nil {
		return fds.Add(fd)
	}

	if pos < 0 {
		pos = 0
	}

	mid := fds
	for mid.pos != pos {
		mid = mid.next
	}

	bot := mid.next
	mid.next = nil
	fds.pos -= mid.pos

	bot = bot.Add(fd)
	fds.Join(bot)

	return fds
}

// Join attaches the tail of the receiver list "fds" to the head of the otherList.
func (fds *FieldDefinitions) Join(otherList *FieldDefinitions) {
	if fds == nil {
		return
	}

	pos := fds.Len() + otherList.Len() - 1

	last := fds
	for fds != nil {
		fds.pos = pos
		pos--
		last = fds
		fds = fds.next
	}

	last.next = otherList
}

// Len returns the length of this linked list.
func (fds *FieldDefinitions) Len() int {
	if fds == nil {
		return 0
	}
	return fds.pos + 1
}

// Reverse reverses this linked list of FieldDefinition. Usually when the linked list is being
// constructed the result will be last-to-first, so we'll want to reverse it to get it in the
// "right" order.
func (fds *FieldDefinitions) Reverse() *FieldDefinitions {
	current := fds

	var prev *FieldDefinitions
	var pos int

	for current != nil {
		current.pos = pos
		pos++

		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

// FieldDefinitionsFromSlice returns a FieldDefinitions list from a slice of FieldDefinition.
func FieldDefinitionsFromSlice(sl []FieldDefinition) *FieldDefinitions {
	var list *FieldDefinitions
	for _, v := range sl {
		list = list.Add(v)
	}
	return list.Reverse()
}

// InputValueDefinitions is a linked list that contains InputValueDefinition values.
type InputValueDefinitions struct {
	Data InputValueDefinition
	next *InputValueDefinitions
	pos  int
}

// Add appends a InputValueDefinition to this linked list and returns this new head.
func (ivds *InputValueDefinitions) Add(data InputValueDefinition) *InputValueDefinitions {
	var pos int

	if ivds != nil {
		pos = ivds.pos + 1
	}

	return &InputValueDefinitions{
		Data: data,
		next: ivds,
		pos:  pos,
	}
}

// ForEach applies the given map function to each item in this linked list.
func (ivds *InputValueDefinitions) ForEach(fn func(ivd InputValueDefinition, i int)) {
	if ivds == nil {
		return
	}

	iter := 0
	current := ivds

	for {
		fn(current.Data, iter)

		if current.next == nil {
			break
		}

		iter++
		current = current.next
	}
}

// Insert places the InputValueDefinition in the position given by pos.
// The method will insert at top if pos is greater than or equal to list length.
// The method will insert at bottom if the pos is less than 0.
func (ivds *InputValueDefinitions) Insert(ivd InputValueDefinition, pos int) *InputValueDefinitions {
	if pos >= ivds.Len() || ivds == nil {
		return ivds.Add(ivd)
	}

	if pos < 0 {
		pos = 0
	}

	mid := ivds
	for mid.pos != pos {
		mid = mid.next
	}

	bot := mid.next
	mid.next = nil
	ivds.pos -= mid.pos

	bot = bot.Add(ivd)
	ivds.Join(bot)

	return ivds
}

// Join attaches the tail of the receiver list "ivds" to the head of the otherList.
func (ivds *InputValueDefinitions) Join(otherList *InputValueDefinitions) {
	if ivds == nil {
		return
	}

	pos := ivds.Len() + otherList.Len() - 1

	last := ivds
	for ivds != nil {
		ivds.pos = pos
		pos--
		last = ivds
		ivds = ivds.next
	}

	last.next = otherList
}

// Len returns the length of this linked list.
func (ivds *InputValueDefinitions) Len() int {
	if ivds == nil {
		return 0
	}
	return ivds.pos + 1
}

// Reverse reverses this linked list of InputValueDefinition. Usually when the linked list is being
// constructed the result will be last-to-first, so we'll want to reverse it to get it in the
// "right" order.
func (ivds *InputValueDefinitions) Reverse() *InputValueDefinitions {
	current := ivds

	var prev *InputValueDefinitions
	var pos int

	for current != nil {
		current.pos = pos
		pos++

		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

// InputValueDefinitionsFromSlice returns a InputValueDefinitions list from a slice of InputValueDefinition.
func InputValueDefinitionsFromSlice(sl []InputValueDefinition) *InputValueDefinitions {
	var list *InputValueDefinitions
	for _, v := range sl {
		list = list.Add(v)
	}
	return list.Reverse()
}

// OperationTypeDefinitions is a linked list that contains OperationTypeDefinition values.
type OperationTypeDefinitions struct {
	Data OperationTypeDefinition
	next *OperationTypeDefinitions
	pos  int
}

// Add appends a OperationTypeDefinition to this linked list and returns this new head.
func (otds *OperationTypeDefinitions) Add(data OperationTypeDefinition) *OperationTypeDefinitions {
	var pos int

	if otds != nil {
		pos = otds.pos + 1
	}

	return &OperationTypeDefinitions{
		Data: data,
		next: otds,
		pos:  pos,
	}
}

// ForEach applies the given map function to each item in this linked list.
func (otds *OperationTypeDefinitions) ForEach(fn func(otd OperationTypeDefinition, i int)) {
	if otds == nil {
		return
	}

	iter := 0
	current := otds

	for {
		fn(current.Data, iter)

		if current.next == nil {
			break
		}

		iter++
		current = current.next
	}
}

// Insert places the OperationTypeDefinition in the position given by pos.
// The method will insert at top if pos is greater than or equal to list length.
// The method will insert at bottom if the pos is less than 0.
func (otds *OperationTypeDefinitions) Insert(otd OperationTypeDefinition, pos int) *OperationTypeDefinitions {
	if pos >= otds.Len() || otds == nil {
		return otds.Add(otd)
	}

	if pos < 0 {
		pos = 0
	}

	mid := otds
	for mid.pos != pos {
		mid = mid.next
	}

	bot := mid.next
	mid.next = nil
	otds.pos -= mid.pos

	bot = bot.Add(otd)
	otds.Join(bot)

	return otds
}

// Join attaches the tail of the receiver list "otds" to the head of the otherList.
func (otds *OperationTypeDefinitions) Join(otherList *OperationTypeDefinitions) {
	if otds == nil {
		return
	}

	pos := otds.Len() + otherList.Len() - 1

	last := otds
	for otds != nil {
		otds.pos = pos
		pos--
		last = otds
		otds = otds.next
	}

	last.next = otherList
}

// Len returns the length of this linked list.
func (otds *OperationTypeDefinitions) Len() int {
	if otds == nil {
		return 0
	}
	return otds.pos + 1
}

// Reverse reverses this linked list of OperationTypeDefinition. Usually when the linked list is being
// constructed the result will be last-to-first, so we'll want to reverse it to get it in the
// "right" order.
func (otds *OperationTypeDefinitions) Reverse() *OperationTypeDefinitions {
	current := otds

	var prev *OperationTypeDefinitions
	var pos int

	for current != nil {
		current.pos = pos
		pos++

		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

// OperationTypeDefinitionsFromSlice returns a OperationTypeDefinitions list from a slice of OperationTypeDefinition.
func OperationTypeDefinitionsFromSlice(sl []OperationTypeDefinition) *OperationTypeDefinitions {
	var list *OperationTypeDefinitions
	for _, v := range sl {
		list = list.Add(v)
	}
	return list.Reverse()
}

// RootOperationTypeDefinitions is a linked list that contains RootOperationTypeDefinition values.
type RootOperationTypeDefinitions struct {
	Data RootOperationTypeDefinition
	next *RootOperationTypeDefinitions
	pos  int
}

// Add appends a RootOperationTypeDefinition to this linked list and returns this new head.
func (rotds *RootOperationTypeDefinitions) Add(data RootOperationTypeDefinition) *RootOperationTypeDefinitions {
	var pos int

	if rotds != nil {
		pos = rotds.pos + 1
	}

	return &RootOperationTypeDefinitions{
		Data: data,
		next: rotds,
		pos:  pos,
	}
}

// ForEach applies the given map function to each item in this linked list.
func (rotds *RootOperationTypeDefinitions) ForEach(fn func(rotd RootOperationTypeDefinition, i int)) {
	if rotds == nil {
		return
	}

	iter := 0
	current := rotds

	for {
		fn(current.Data, iter)

		if current.next == nil {
			break
		}

		iter++
		current = current.next
	}
}

// Insert places the RootOperationTypeDefinition in the position given by pos.
// The method will insert at top if pos is greater than or equal to list length.
// The method will insert at bottom if the pos is less than 0.
func (rotds *RootOperationTypeDefinitions) Insert(rotd RootOperationTypeDefinition, pos int) *RootOperationTypeDefinitions {
	if pos >= rotds.Len() || rotds == nil {
		return rotds.Add(rotd)
	}

	if pos < 0 {
		pos = 0
	}

	mid := rotds
	for mid.pos != pos {
		mid = mid.next
	}

	bot := mid.next
	mid.next = nil
	rotds.pos -= mid.pos

	bot = bot.Add(rotd)
	rotds.Join(bot)

	return rotds
}

// Join attaches the tail of the receiver list "rotds" to the head of the otherList.
func (rotds *RootOperationTypeDefinitions) Join(otherList *RootOperationTypeDefinitions) {
	if rotds == nil {
		return
	}

	pos := rotds.Len() + otherList.Len() - 1

	last := rotds
	for rotds != nil {
		rotds.pos = pos
		pos--
		last = rotds
		rotds = rotds.next
	}

	last.next = otherList
}

// Len returns the length of this linked list.
func (rotds *RootOperationTypeDefinitions) Len() int {
	if rotds == nil {
		return 0
	}
	return rotds.pos + 1
}

// Reverse reverses this linked list of RootOperationTypeDefinition. Usually when the linked list is being
// constructed the result will be last-to-first, so we'll want to reverse it to get it in the
// "right" order.
func (rotds *RootOperationTypeDefinitions) Reverse() *RootOperationTypeDefinitions {
	current := rotds

	var prev *RootOperationTypeDefinitions
	var pos int

	for current != nil {
		current.pos = pos
		pos++

		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

// RootOperationTypeDefinitionsFromSlice returns a RootOperationTypeDefinitions list from a slice of RootOperationTypeDefinition.
func RootOperationTypeDefinitionsFromSlice(sl []RootOperationTypeDefinition) *RootOperationTypeDefinitions {
	var list *RootOperationTypeDefinitions
	for _, v := range sl {
		list = list.Add(v)
	}
	return list.Reverse()
}

// Selections is a linked list that contains Selection values.
type Selections struct {
	Data Selection
	next *Selections
	pos  int
}

// Add appends a Selection to this linked list and returns this new head.
func (ss *Selections) Add(data Selection) *Selections {
	var pos int

	if ss != nil {
		pos = ss.pos + 1
	}

	return &Selections{
		Data: data,
		next: ss,
		pos:  pos,
	}
}

// ForEach applies the given map function to each item in this linked list.
func (ss *Selections) ForEach(fn func(s Selection, i int)) {
	if ss == nil {
		return
	}

	iter := 0
	current := ss

	for {
		fn(current.Data, iter)

		if current.next == nil {
			break
		}

		iter++
		current = current.next
	}
}

// Insert places the Selection in the position given by pos.
// The method will insert at top if pos is greater than or equal to list length.
// The method will insert at bottom if the pos is less than 0.
func (ss *Selections) Insert(s Selection, pos int) *Selections {
	if pos >= ss.Len() || ss == nil {
		return ss.Add(s)
	}

	if pos < 0 {
		pos = 0
	}

	mid := ss
	for mid.pos != pos {
		mid = mid.next
	}

	bot := mid.next
	mid.next = nil
	ss.pos -= mid.pos

	bot = bot.Add(s)
	ss.Join(bot)

	return ss
}

// Join attaches the tail of the receiver list "ss" to the head of the otherList.
func (ss *Selections) Join(otherList *Selections) {
	if ss == nil {
		return
	}

	pos := ss.Len() + otherList.Len() - 1

	last := ss
	for ss != nil {
		ss.pos = pos
		pos--
		last = ss
		ss = ss.next
	}

	last.next = otherList
}

// Len returns the length of this linked list.
func (ss *Selections) Len() int {
	if ss == nil {
		return 0
	}
	return ss.pos + 1
}

// Reverse reverses this linked list of Selection. Usually when the linked list is being
// constructed the result will be last-to-first, so we'll want to reverse it to get it in the
// "right" order.
func (ss *Selections) Reverse() *Selections {
	current := ss

	var prev *Selections
	var pos int

	for current != nil {
		current.pos = pos
		pos++

		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

// SelectionsFromSlice returns a Selections list from a slice of Selection.
func SelectionsFromSlice(sl []Selection) *Selections {
	var list *Selections
	for _, v := range sl {
		list = list.Add(v)
	}
	return list.Reverse()
}

// Types is a linked list that contains Type values.
type Types struct {
	Data Type
	next *Types
	pos  int
}

// Add appends a Type to this linked list and returns this new head.
func (ts *Types) Add(data Type) *Types {
	var pos int

	if ts != nil {
		pos = ts.pos + 1
	}

	return &Types{
		Data: data,
		next: ts,
		pos:  pos,
	}
}

// ForEach applies the given map function to each item in this linked list.
func (ts *Types) ForEach(fn func(t Type, i int)) {
	if ts == nil {
		return
	}

	iter := 0
	current := ts

	for {
		fn(current.Data, iter)

		if current.next == nil {
			break
		}

		iter++
		current = current.next
	}
}

// Insert places the Type in the position given by pos.
// The method will insert at top if pos is greater than or equal to list length.
// The method will insert at bottom if the pos is less than 0.
func (ts *Types) Insert(t Type, pos int) *Types {
	if pos >= ts.Len() || ts == nil {
		return ts.Add(t)
	}

	if pos < 0 {
		pos = 0
	}

	mid := ts
	for mid.pos != pos {
		mid = mid.next
	}

	bot := mid.next
	mid.next = nil
	ts.pos -= mid.pos

	bot = bot.Add(t)
	ts.Join(bot)

	return ts
}

// Join attaches the tail of the receiver list "ts" to the head of the otherList.
func (ts *Types) Join(otherList *Types) {
	if ts == nil {
		return
	}

	pos := ts.Len() + otherList.Len() - 1

	last := ts
	for ts != nil {
		ts.pos = pos
		pos--
		last = ts
		ts = ts.next
	}

	last.next = otherList
}

// Len returns the length of this linked list.
func (ts *Types) Len() int {
	if ts == nil {
		return 0
	}
	return ts.pos + 1
}

// Reverse reverses this linked list of Type. Usually when the linked list is being
// constructed the result will be last-to-first, so we'll want to reverse it to get it in the
// "right" order.
func (ts *Types) Reverse() *Types {
	current := ts

	var prev *Types
	var pos int

	for current != nil {
		current.pos = pos
		pos++

		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

// TypesFromSlice returns a Types list from a slice of Type.
func TypesFromSlice(sl []Type) *Types {
	var list *Types
	for _, v := range sl {
		list = list.Add(v)
	}
	return list.Reverse()
}

// VariableDefinitions is a linked list that contains VariableDefinition values.
type VariableDefinitions struct {
	Data VariableDefinition
	next *VariableDefinitions
	pos  int
}

// Add appends a VariableDefinition to this linked list and returns this new head.
func (vds *VariableDefinitions) Add(data VariableDefinition) *VariableDefinitions {
	var pos int

	if vds != nil {
		pos = vds.pos + 1
	}

	return &VariableDefinitions{
		Data: data,
		next: vds,
		pos:  pos,
	}
}

// ForEach applies the given map function to each item in this linked list.
func (vds *VariableDefinitions) ForEach(fn func(vd VariableDefinition, i int)) {
	if vds == nil {
		return
	}

	iter := 0
	current := vds

	for {
		fn(current.Data, iter)

		if current.next == nil {
			break
		}

		iter++
		current = current.next
	}
}

// Insert places the VariableDefinition in the position given by pos.
// The method will insert at top if pos is greater than or equal to list length.
// The method will insert at bottom if the pos is less than 0.
func (vds *VariableDefinitions) Insert(vd VariableDefinition, pos int) *VariableDefinitions {
	if pos >= vds.Len() || vds == nil {
		return vds.Add(vd)
	}

	if pos < 0 {
		pos = 0
	}

	mid := vds
	for mid.pos != pos {
		mid = mid.next
	}

	bot := mid.next
	mid.next = nil
	vds.pos -= mid.pos

	bot = bot.Add(vd)
	vds.Join(bot)

	return vds
}

// Join attaches the tail of the receiver list "vds" to the head of the otherList.
func (vds *VariableDefinitions) Join(otherList *VariableDefinitions) {
	if vds == nil {
		return
	}

	pos := vds.Len() + otherList.Len() - 1

	last := vds
	for vds != nil {
		vds.pos = pos
		pos--
		last = vds
		vds = vds.next
	}

	last.next = otherList
}

// Len returns the length of this linked list.
func (vds *VariableDefinitions) Len() int {
	if vds == nil {
		return 0
	}
	return vds.pos + 1
}

// Reverse reverses this linked list of VariableDefinition. Usually when the linked list is being
// constructed the result will be last-to-first, so we'll want to reverse it to get it in the
// "right" order.
func (vds *VariableDefinitions) Reverse() *VariableDefinitions {
	current := vds

	var prev *VariableDefinitions
	var pos int

	for current != nil {
		current.pos = pos
		pos++

		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

// VariableDefinitionsFromSlice returns a VariableDefinitions list from a slice of VariableDefinition.
func VariableDefinitionsFromSlice(sl []VariableDefinition) *VariableDefinitions {
	var list *VariableDefinitions
	for _, v := range sl {
		list = list.Add(v)
	}
	return list.Reverse()
}
