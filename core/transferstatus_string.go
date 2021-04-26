// Code generated by "stringer -type TransferStatus -trimprefix TransferStatus"; DO NOT EDIT.

package core

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TransferStatusPending-0]
	_ = x[TransferStatusAssigned-1]
	_ = x[TransferStatusHandled-2]
	_ = x[TransferStatusPassed-3]
}

const _TransferStatus_name = "PendingAssignedHandledPassed"

var _TransferStatus_index = [...]uint8{0, 7, 15, 22, 28}

func (i TransferStatus) String() string {
	if i < 0 || i >= TransferStatus(len(_TransferStatus_index)-1) {
		return "TransferStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TransferStatus_name[_TransferStatus_index[i]:_TransferStatus_index[i+1]]
}
