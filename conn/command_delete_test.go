package conn_test

import (
	"github.com/jordwest/imap-server/conn"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("DELETE Command", func() {
	Context("When not logged in", func() {
		BeforeEach(func() {
			tConn.SetState(conn.StateNotAuthenticated)
		})

		It("should give an error", func() {
			SendLine(`abcd.123 DELETE "INBOX"`)
			ExpectResponse("abcd.123 BAD not authenticated")
		})
	})

	Context("When logged in", func() {
		BeforeEach(func() {
			tConn.SetState(conn.StateAuthenticated)
			tConn.User = mStore.User
		})

		It("delete folder should work without errors", func() {
			SendLine(`abcd.123 DELETE "INBOX"`)
			ExpectResponse("abcd.123 OK DELETE Completed")
		})

		It("delete folder with invalid name give an error", func() {
			SendLine(`abcd.123 DELETE "invalid"`)
			ExpectResponse("abcd.123 NO DELETE Failure: can't delete mailbox with that name")
		})
	})
})
