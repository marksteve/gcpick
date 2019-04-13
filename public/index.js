/* global fetch List */

fetch('/projects')
  .then(r => r.json())
  .then(createList)

function createList (projects) {
  const options = {
    item: 'project',
    valueNames: ['name', 'projectId']
  }
  const list = new List('projects', options, projects)
  list.sort('createTime', { order: 'desc' })
  list.on('searchComplete', function (e) {
    const found = e.matchingItems.length === 1
    toggleLinks(found)
    if (found) {
      updateLinks(e.matchingItems[0])
    }
  })
}

const links = document.querySelector('.links')

function toggleLinks (isToggled) {
  links.classList.toggle('-visible', isToggled)
}

function updateLinks (match) {
  const { projectId } = match.values()
  Array.from(links.querySelectorAll('a')).map(function (a) {
    a.setAttribute('href', a.dataset.href + projectId)
  })
}
