# svn-files

[![Build Status](https://travis-ci.org/tanelpuhu/svn-files.svg?branch=master)](https://travis-ci.org/tanelpuhu/svn-files)
[![Go Report Card](https://goreportcard.com/badge/github.com/tanelpuhu/svn-files)](https://goreportcard.com/report/github.com/tanelpuhu/svn-files)

Usage:

	svn-files [path-to-direcory]

Example:

	$ svnadmin create repo
	$ svn co file://$PWD/repo checkout && cd checkout
	Checked out revision 0.
	$ touch yks kaks kolm && svn add * && svn ci -m init
	A         kaks
	A         kolm
	A         yks
	Adding         kaks
	Adding         kolm
	Adding         yks
	Transmitting file data ...done
	Committing transaction...
	Committed revision 1.
	$ date >> kolm && svn ci -m "date to kolm"
	Sending        kolm
	Transmitting file data .done
	Committing transaction...
	Committed revision 2.
	$ date >> yks && svn ci -m "date to yks"
	Sending        yks
	Transmitting file data .done
	Committing transaction...
	Committed revision 3.
	$ svn up
	Updating '.':
	At revision 3.
	$ svn log -qv
	------------------------------------------------------------------------
	r3 | tanel | 2018-04-06 23:41:17 +0300 (Fri, 06 Apr 2018)
	Changed paths:
	   M /yks
	------------------------------------------------------------------------
	r2 | tanel | 2018-04-06 23:41:15 +0300 (Fri, 06 Apr 2018)
	Changed paths:
	   M /kolm
	------------------------------------------------------------------------
	r1 | tanel | 2018-04-06 23:41:13 +0300 (Fri, 06 Apr 2018)
	Changed paths:
	   A /kaks
	   A /kolm
	   A /yks
	------------------------------------------------------------------------
	$ svn-files
	2018-04-06 20:41:13  1  tanel  kaks
	2018-04-06 20:41:15  2  tanel  kolm
	2018-04-06 20:41:17  3  tanel  yks
	$ date >> yks && svn ci -m "date to yks"
	Sending        yks
	Transmitting file data .done
	Committing transaction...
	Committed revision 4.
	$ svn-files
	2018-04-06 20:41:13  1  tanel  kaks
	2018-04-06 20:41:15  2  tanel  kolm
	2018-04-06 20:41:17  3  tanel  yks
	$ svn up
	Updating '.':
	At revision 4.
	$ svn-files
	2018-04-06 20:41:13  1  tanel  kaks
	2018-04-06 20:41:15  2  tanel  kolm
	2018-04-06 20:41:38  4  tanel  yks
	$ date >> kaks && svn ci -m "date to kaks"
	Sending        kaks
	Transmitting file data .done
	Committing transaction...
	Committed revision 5.
	$ svn up
	Updating '.':
	At revision 5.
	$ svn log -qv
	------------------------------------------------------------------------
	r5 | tanel | 2018-04-06 23:42:01 +0300 (Fri, 06 Apr 2018)
	Changed paths:
	   M /kaks
	------------------------------------------------------------------------
	r4 | tanel | 2018-04-06 23:41:38 +0300 (Fri, 06 Apr 2018)
	Changed paths:
	   M /yks
	------------------------------------------------------------------------
	r3 | tanel | 2018-04-06 23:41:17 +0300 (Fri, 06 Apr 2018)
	Changed paths:
	   M /yks
	------------------------------------------------------------------------
	r2 | tanel | 2018-04-06 23:41:15 +0300 (Fri, 06 Apr 2018)
	Changed paths:
	   M /kolm
	------------------------------------------------------------------------
	r1 | tanel | 2018-04-06 23:41:13 +0300 (Fri, 06 Apr 2018)
	Changed paths:
	   A /kaks
	   A /kolm
	   A /yks
	------------------------------------------------------------------------
	$ svn-files
	2018-04-06 20:41:15  2  tanel  kolm
	2018-04-06 20:41:38  4  tanel  yks
	2018-04-06 20:42:01  5  tanel  kaks

