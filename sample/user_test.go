// A test that uses a mock.
package user_test

import (
	"testing"

	"gomock.googlecode.com/hg/gomock"
	"gomock.googlecode.com/hg/sample/mock_user"
	"gomock.googlecode.com/hg/sample/user"
)

func TestRemember(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIndex := mock_user.NewMockIndex(ctrl)
	mockIndex.EXPECT().Put("a", 1)            // literals work
	mockIndex.EXPECT().Put("b", gomock.Eq(2)) // matchers work too

	// NillableRet returns os.Error. Not declaring it should result in a nil return.
	mockIndex.EXPECT().NillableRet()

	// Should be able to place expectations on variadic methods.
	mockIndex.EXPECT().Ellip("%d", 0, 1, 1, 2, 3) // direct args
	tri := []interface{}{1, 3, 6, 10, 15}
	mockIndex.EXPECT().Ellip("%d", tri...) // args from slice

	user.Remember(mockIndex, []string{"a", "b"}, []interface{}{1, 2})

	// Try one with an action.
	calledString := ""
	mockIndex.EXPECT().Put(gomock.Any(), gomock.Any()).Do(func(key string, _ interface{}) {
		calledString = key
	})
	mockIndex.EXPECT().NillableRet()
	user.Remember(mockIndex, []string{"blah"}, []interface{}{7})
	if calledString != "blah" {
		t.Fatalf(`Uh oh. %q != "blah"`, calledString)
	}
}

func TestGrabPointer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIndex := mock_user.NewMockIndex(ctrl)
	mockIndex.EXPECT().Ptr(gomock.Any()).SetArg(0, 7) // set first argument to 7

	i := user.GrabPointer(mockIndex)
	if i != 7 {
		t.Errorf("Expected 7, got %d", i)
	}
}
