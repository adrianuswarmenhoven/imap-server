package conn_test

import (
	"github.com/adrianuswarmenhoven/imap-server/conn"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("CHECK Command", func() {
	Context("When logged in", func() {
		BeforeEach(func() {
			tConn.SetState(conn.StateAuthenticated)
			tConn.User = mStore.User
		})

		It("shoud send response", func() {
			SendLine(`abcd.123 CHECK`)
			ExpectResponse("abcd.123 OK CHECK Completed")
		})
	})

	Context("When not logged in", func() {
		BeforeEach(func() {
			tConn.SetState(conn.StateNotAuthenticated)
		})

		It("shoud send response", func() {
			SendLine(`abcd.123 CHECK`)
			ExpectResponse("abcd.123 OK CHECK Completed")
		})
	})
})
