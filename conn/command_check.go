package conn

func cmdCheck(args commandArgs, c *Conn) {
	c.writeResponse(args.ID(), "OK CHECK Completed")
}
