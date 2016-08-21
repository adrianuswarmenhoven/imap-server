package conn_test

import (
	"github.com/jordwest/imap-server/conn"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("SUBSCRIBE Command", func() {
	Context("When not logged in", func() {
		BeforeEach(func() {
			tConn.SetState(conn.StateNotAuthenticated)
		})

		It("should give an error", func() {
			SendLine(`abcd.123 SUBSCRIBE "test"`)
			ExpectResponse("abcd.123 BAD not authenticated")
		})
	})

	Context("When logged in", func() {
		BeforeEach(func() {
			tConn.SetState(conn.StateAuthenticated)
			tConn.User = mStore.User
		})

		It("should work without errors", func() {
			SendLine(`abcd.123 SUBSCRIBE "test"`)
			ExpectResponse("abcd.123 OK SUBSCRIBE completed")
		})
	})
})
