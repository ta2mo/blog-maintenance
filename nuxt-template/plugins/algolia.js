import Algoliasearch from 'algoliasearch'

const client = Algoliasearch('WT16MPP7YO', '4805b40bc559a8a08bf9909cf56d033e');
const index = client.initIndex('ta2mo.github.io');

export default index
