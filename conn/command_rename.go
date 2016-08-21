package conn

func cmdRename(args commandArgs, c *Conn) {
	if !c.assertAuthenticated(args.ID()) {
		return
	}
	if err := c.User.RenameMailbox(args.Arg(0), args.Arg(1)); err != nil {
		c.writeResponse(args.ID(), "NO RENAME Failure: can't rename mailbox with that name, can't rename to mailbox with that name")
		return
	}
	c.writeResponse(args.ID(), "OK RENAME Completed")
}
