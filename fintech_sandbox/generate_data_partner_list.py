#!/usr/bin/env python2
# -*- coding: utf-8 -*-
"""
Created on Tue Jan 24 16:52:08 2017

@author: Sean Kruzel
"""

from bs4 import BeautifulSoup
import requests
import pandas as pd
import urlparse

WIKI_REPO_DIR = '/home/closedloop/GitHub/consulting/fintech-sandbox-curation.wiki/'

def naming_convention(full_name):
    import re
    N = len(full_name)
    full_name = full_name.lower().strip().replace(' ','_', N).replace(',','', N).replace('.','_', N)
    full_name = re.sub('__*', '_', full_name) # remove duplicate underscores
    full_name = re.sub('_$', '', full_name) # remove leading underscore
    full_name = re.sub('^_', '', full_name) # remove trailing underscore
    return full_name
  
def process_article(html):
    import shutil
    images_dir = WIKI_REPO_DIR + "images/"
    
    soup = BeautifulSoup(html, 'lxml')
    
    elts = soup.findAll('article')
    profiles = []
    for elt in elts:
        a = elt.find('a')
        img = elt.find('img')
        text = elt.get_text().strip().splitlines()
        if len(text):
            text = text[0]
            
        logo = img.attrs.get('src', None)
        if logo:
            img_name = urlparse.urlparse(logo).path.split('/')[-1]
        else:
            img_name = None

        profile = { 
            'logo': logo,
            'img_name': img_name,
            'id': naming_convention(text),
            'logo_type': img_name.split('.')[-1],
            'name': text,
            'link': a.attrs.get('href')
        }
        if profile['id'][-1] == '_':
            profile['id'] = profile['id'][:-1]
        profiles.append(profile)
    
    for profile in profiles:
        # Process bio
        rs = requests.get('http://fintechsandbox.org' + profile['link'])
        html = rs.content
        soup = BeautifulSoup(html, 'lxml')
        
        elts = soup.findAll('article')
        title_elt = soup.find(name='h1', attrs={'class':'title'})
        if title_elt:
            profile.update({'title':title_elt.text.strip()})
        desc_elt = soup.find(name='div', attrs={'class':'content'})
        if desc_elt:
            profile.update({'title':desc_elt.text.strip().splitlines()[0]})
            
        # Process External Links
        sidebar_elt = soup.find(name='div', attrs={'class':'sidebar'})
        profile.update({'external_links': [{'url':elt.attrs['href'].strip(), 'text': elt.text.strip() } for elt in sidebar_elt.findAll('a')]})
        
        for l in profile['external_links']:
            if l['url'].find('mailto:') == 0:
                profile['email_link'] = l['url']
                profile['email_text'] = l['text']
            elif l['url'].find('tel:') == 0:
                profile['phone_link'] = l['url']
                profile['phone_text'] = l['text']
            elif l['url'].find('linkedin') > -1:
                profile['linkedin'] = l['url']
            elif l['url'].find('twitter') > -1:
                profile['twitter'] = l['url']
            elif l['url'].find('fintechsandbox') > -1:
                pass
            else:
                profile['website'] = l['url']
                
    
        # Download logo
        rs = requests.get(profile['logo'], stream=True)
        logo_fname = '{}{}.{}'.format(images_dir, profile['id'], profile['logo_type'])
        profile.update({'logo_fname': logo_fname })
        profile['logo'] = profile['logo_fname'].split('/')[-1]
        if rs.status_code == 200:
            with open(logo_fname, 'wb') as f:
                rs.raw.decode_content = True
                shutil.copyfileobj(rs.raw, f)
        
    import pandas as pd
    df = pd.DataFrame(profiles)
    df['name'] = df['name'].apply(lambda x: x.strip())
    return df

    
def bootstrap_members():
    """ Pull in information for each profile
    
    Required Fields:
        id
        Full Name
        website
        logo
        description
        category
        sandbox_start_date
    Optional Fields:
        twitter
        linkedin
        crunchbase
        instagram
        facebook
        company linked-in profile
        sandbox_alumni_start_date
    Detailed Fields
        Employees
            name
            email
            twitter
            crunchbase
            linkedin
            Other social
                facebook
                instagram
                snapchat
                other ...
            roles: (Founder, Sales Contact, Tech Contact)
            is_fintech_sandbox_contact        

    """
    import requests

    url = "http://fintechsandbox.org/startups"

    rs = requests.get(url)
    if rs.status_code == 200:
        html = rs.content
        return process_article(html)
    else:
        raise ValueError, "url: {} - is not available".format(url)
    

def bootstrap_partners():
    """ Pull in information for each partner
    
    Required Fields:
        id
        Full Name
        website
        logo
        description
        category
        sandbox_start_date
    Optional Fields:
        twitter
        linkedin
        crunchbase
        instagram
        facebook
        company linked-in profile
        sandbox_alumni_start_date
    Detailed Fields
        Employees
            name
            email
            twitter
            crunchbase
            linkedin
            Other social
                facebook
                instagram
                snapchat
                other ...
            roles: (Founder, Sales Contact, Tech Contact)
            is_fintech_sandbox_contact        

    """
    import requests
    
    url = "http://fintechsandbox.org/data-partners"
    rs = requests.get(url)
    if rs.status_code == 200:
        html = rs.content
        return process_article(html)
    else:
        raise ValueError, "url: {} - is not available".format(url)
    

def create_markup_table(df, profile_category):
    """Given a DataFrame where each row is an entity, and each column is a field
    map the fields to text in columns in the table"""
    
    assert profile_category in ['Sandbox-Members','Data-Providers'], "profile_category must be either 'Sandbox-Members' or 'Data-Providers'"
    #[![Fintech Sandbox Profile](images/thomson_reuters.png)](http://fintechsandbox.org/partner/thomson-reuters)
    
    headings = ['Provider',
                'Resource Page', 
                'External Profile', 'Website', 'Email', 'Phone','LinkedIn','Twitter']
    template = ['[[images/{logo}]]',
                '[{name}]({id})',
                '[![Fintech Sandbox Profile](images/icons/external_link.png)](http://fintechsandbox.org{link})',
                '[![Website](images/icons/www.png)]({website})',
                '[![{email_text}](images/icons/email.png)]({email_link})',
                '[![{phone_text}](images/icons/phone.png)]({phone_link})',
                '[![{linkedin}](images/icons/linkedin.png)]({linkedin})',
                '[![{twitter}](images/icons/twitter.png)]({twitter})'
    ]
    
    def _is_nan(v):
        import numpy as np
        if v is None:
            return True
        else:
            try:
                return np.isnan(v)
            except:
                return False
    
    def _process_template(template_string, data):
        
        try:
            props = { k:v for k,v in data.to_dict().items() if not _is_nan(v) }
            return t.format(**props)
        except KeyError:
            return ''        
    
    table = ['| ' + ' | '.join(headings) +' |']
    table.append('|' + '|'.join(['-'*len(h) for h in headings]) +'|')
    for _, data in df.iterrows():                
        row = '| ' + ' | '.join([_process_template(t,data) for t in template]) +' |'
        table.append(row)

    markup = '\n'.join(table)
    
    with open(WIKI_REPO_DIR + '{}.md'.format(profile_category), 'w') as fh:
        fh.writelines(markup)
    
def create_profile(profile, profile_category=None):
    assert profile_category in ['Sandbox-Members','Data-Providers'], "profile_category must be either 'Sandbox-Members' or 'Data-Providers'"
    raise NotImplementedError
    
    
if __name__ == "__main__":
    
    BOOTSTRAP = True
    REFRESH_TABLE_OF_CONTENTS = True
    CREATE_PROFILES = False
    if BOOTSTRAP:
        print "Loading content from FinTech Sandbox Website"

        print "Loading members"
        members = bootstrap_members()
        members.to_csv(WIKI_REPO_DIR + 'members.csv', encoding='utf-8')
        print "{} members info are loaded".format(len(members))

        print "Loading Providers"
        providers = bootstrap_partners()
        providers.to_csv(WIKI_REPO_DIR + 'providers.csv', encoding='utf-8')
        print "{} providers info are loaded".format(len(providers))
    else:
        members = pd.read_csv(WIKI_REPO_DIR + 'members.csv')
        providers = pd.read_csv(WIKI_REPO_DIR + 'providers.csv')

    if REFRESH_TABLE_OF_CONTENTS:
        create_markup_table(members, profile_category = 'Sandbox-Members')
        create_markup_table(providers, profile_category = 'Data-Providers')
    
    if CREATE_PROFILES:
        for _, profile in members:
            create_profile(profile, profile_category= 'Sandbox-Members')

        for _, profile in providers:
            create_profile(profile, profile_category= 'Data-Providers')
