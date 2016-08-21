package conn_test

import (
	"github.com/jordwest/imap-server/conn"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Create Command", func() {
	Context("When not logged in", func() {
		BeforeEach(func() {
			tConn.SetState(conn.StateNotAuthenticated)
		})

		It("should give an error", func() {
			SendLine(`abcd.123 CREATE "test"`)
			ExpectResponse("abcd.123 BAD not authenticated")
		})
	})

	Context("When logged in", func() {
		BeforeEach(func() {
			tConn.SetState(conn.StateAuthenticated)
			tConn.User = mStore.User
		})

		It("should work without errors", func() {
			SendLine(`abcd.123 CREATE "test"`)
			ExpectResponse("abcd.123 OK CREATE Completed")
		})

		It("invalid name give an error", func() {
			SendLine(`abcd.123 CREATE "invalid"`)
			ExpectResponse("abcd.123 NO create failure: can't create mailbox with that name")
		})
	})
})
