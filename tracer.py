#!/usr/bin/env python

from SimpleHTTPServer import SimpleHTTPRequestHandler
from BaseHTTPServer import HTTPServer
from urlparse import urlparse, parse_qs
import cgi

PORT=5003

class TracerHandler(SimpleHTTPRequestHandler):
    
    container={'k':0};
    
    def do_GET(self):
        #self.send_response(200)
        
        ctype, pdict = cgi.parse_header(self.headers.getheader('content-type'))
        if ctype == 'multipart/form-data':
            postvars = cgi.parse_multipart(self.rfile, pdict)
        elif ctype == 'application/x-www-form-urlencoded':
            length = int(self.headers.getheader('content-length'))
            postvars = cgi.parse_qs(self.rfile.read(length), keep_blank_values=1)
        else:
            postvars = {}

        if 'trace' in postvars.keys():

            self.container['k'] += 1
            
            print str(self.container['k'])+' - '+str(postvars['name'].pop())
            print '---------------------'
            print '\n'
            print postvars['trace'].pop()
            print '\n'
            print '---------------------'
        
        self.wfile.write('Logger')
        
        return
        
    def do_POST(self):
        self.do_GET()
        return
    
        #self.send_response(200)

        
httpd = HTTPServer(("", PORT), TracerHandler)

print "PicPay Tracer", PORT

httpd.serve_forever()