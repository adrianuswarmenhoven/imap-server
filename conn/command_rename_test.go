package conn_test

import (
	"github.com/jordwest/imap-server/conn"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("RENAME Command", func() {
	Context("When not logged in", func() {
		BeforeEach(func() {
			tConn.SetState(conn.StateNotAuthenticated)
		})

		It("should give an error", func() {
			SendLine(`abcd.123 RENAME "old" "new"`)
			ExpectResponse("abcd.123 BAD not authenticated")
		})
	})

	Context("When logged in", func() {
		BeforeEach(func() {
			tConn.SetState(conn.StateAuthenticated)
			tConn.User = mStore.User
		})

		It("should work without errors", func() {
			SendLine(`abcd.123 RENAME "INBOX" "new"`)
			ExpectResponse("abcd.123 OK RENAME Completed")
		})

		It("invalid name give an error", func() {
			SendLine(`abcd.123 RENAME "invalid" "new"`)
			ExpectResponse("abcd.123 NO RENAME Failure: can't rename mailbox with that name, can't rename to mailbox with that name")
		})
	})
})
