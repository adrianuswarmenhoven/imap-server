package conn_test

import (
	"github.com/jordwest/imap-server/conn"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("NOOP Command", func() {
	Context("When logged in", func() {
		BeforeEach(func() {
			tConn.SetState(conn.StateAuthenticated)
			tConn.User = mStore.User
		})

		It("shoud send response", func() {
			SendLine(`abcd.123 NOOP`)
			ExpectResponse("abcd.123 OK NOOP Completed")
		})
	})

	Context("When not logged in", func() {
		BeforeEach(func() {
			tConn.SetState(conn.StateNotAuthenticated)
		})

		It("shoud send response", func() {
			SendLine(`abcd.123 NOOP`)
			ExpectResponse("abcd.123 OK NOOP Completed")
		})
	})
})
