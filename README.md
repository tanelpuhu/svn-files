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
	r3 | tanel | 2019-01-02 23:39:34 +0200 (Wed, 02 Jan 2019)
	Changed paths:
	   M /yks
	------------------------------------------------------------------------
	r2 | tanel | 2019-01-02 23:39:32 +0200 (Wed, 02 Jan 2019)
	Changed paths:
	   M /kolm
	------------------------------------------------------------------------
	r1 | tanel | 2019-01-02 23:39:30 +0200 (Wed, 02 Jan 2019)
	Changed paths:
	   A /kaks
	   A /kolm
	   A /yks
	------------------------------------------------------------------------
	$ svn-files
	2019-01-02 23:39:30  1  tanel  A  kaks
	2019-01-02 23:39:32  2  tanel  M  kolm
	2019-01-02 23:39:34  3  tanel  M  yks
	$ date >> yks && svn ci -m "date to yks"
	Sending        yks
	Transmitting file data .done
	Committing transaction...
	Committed revision 4.
	$ svn-files
	2019-01-02 23:39:30  1  tanel  A  kaks
	2019-01-02 23:39:32  2  tanel  M  kolm
	2019-01-02 23:39:34  3  tanel  M  yks
	$ svn up
	Updating '.':
	At revision 4.
	$ svn-files
	2019-01-02 23:39:30  1  tanel  A  kaks
	2019-01-02 23:39:32  2  tanel  M  kolm
	2019-01-02 23:39:42  4  tanel  M  yks
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
	r5 | tanel | 2019-01-02 23:39:54 +0200 (Wed, 02 Jan 2019)
	Changed paths:
	   M /kaks
	------------------------------------------------------------------------
	r4 | tanel | 2019-01-02 23:39:42 +0200 (Wed, 02 Jan 2019)
	Changed paths:
	   M /yks
	------------------------------------------------------------------------
	r3 | tanel | 2019-01-02 23:39:34 +0200 (Wed, 02 Jan 2019)
	Changed paths:
	   M /yks
	------------------------------------------------------------------------
	r2 | tanel | 2019-01-02 23:39:32 +0200 (Wed, 02 Jan 2019)
	Changed paths:
	   M /kolm
	------------------------------------------------------------------------
	r1 | tanel | 2019-01-02 23:39:30 +0200 (Wed, 02 Jan 2019)
	Changed paths:
	   A /kaks
	   A /kolm
	   A /yks
	------------------------------------------------------------------------
	$ svn-files
	2019-01-02 23:39:32  2  tanel  M  kolm
	2019-01-02 23:39:42  4  tanel  M  yks
	2019-01-02 23:39:54  5  tanel  M  kaks